package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// goroutine对应的函数执行结束了，goroutine就结束
// main函数执行完，由main函数创建的所有goroutine一并结束

var wg sync.WaitGroup // 使用sync.WaitGroup来实现goroutine的同步

// sync.WaitGroup内部维护着一个计数器，计数器的值可以增加和减少
// 例如当我们启动了N 个并发任务时，就调用Add(n)方法将计数器值增加n
// 每个任务完成时通过调用Done()方法将计数器减1
// 通过调用Wait()来等待并发任务执行完，当计数器值为0时，表示所有并发任务已经完成

// sync.WaitGroup是一个结构体，传递的时候要传递指针

func f() {
	defer wg.Done() // goroutine结束就登记-1
	rand.Seed(time.Now().UnixNano())
	r1 := rand.Int() // 生成随机int
	r2 := rand.Intn(10) // 0-10之间取随机数
	fmt.Println(r1, r2)
}

func main() {
	for i :=0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go f()
	}
	wg.Wait() // 等待wg的计数器减为0，即所有登记的goroutine都结束

}
