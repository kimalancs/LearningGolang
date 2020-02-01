package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var notifyCh = make(chan struct{}, 10) // 匿名结构体，空结构体比int更节省空间，int在64位机器上占八个字节，空结构体不占空间，而且也告诉别人这里不需要值

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
		notifyCh <- struct{}{} // 匿名结构体实例化，再加一个花括号
	}

}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 5个任务
	go func() {
		for j := 1; j <= 50; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	// 开启10个goroutine
	for w := 1; w <= 10; w++ {
		go worker(w, jobs, results)
	}

	go func() {
		for i := 0; i < 50; i++ {
			<-notifyCh
		}
		close(results)
	}()
	
	// 输出结果
	for x := range results {
		fmt.Println(x)
	}
}
