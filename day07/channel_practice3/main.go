package main

import (
	"fmt"
	"time"
)

// 通道阻塞，所有的goroutine都在等它，那么程序会发生死锁，会报错
// 如果开启一个后台的goroutine去等它，不会报错，但是会发生goroutine的泄漏，要避免
// goroutine leak，goroutine由于channel的读/写端退出而一直阻塞，导致goroutine一直占用资源，无法退出
// 或goroutine进入死循环，导致资源一直无法释放

// work pool（goroutine池）
// 在工作中我们通常会使用可以指定启动的goroutine数量，worker pool模式，控制goroutine的数量，防止goroutine泄漏和暴涨

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs { 
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// 开启三个goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	// 5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	// 输出结果
	for a := 1; a <= 5; a++ {
		x := <-results
		fmt.Println(x)
	}

}
