package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
单从线程调度讲，Go语言相比起其他语言的优势在于OS线程是由OS内核来调度的
goroutine则是由Go运行时（runtime）自己的调度器调度的
这个调度器使用一个称为m:n调度的技术（复用/调度m个goroutine到n个OS线程）
其一大特点是goroutine的调度是在用户态下完成的， 不涉及内核态与用户态之间的频繁切换
包括内存的分配与释放，都是在用户态维护着一块大的内存池,不直接调用系统的malloc函数（除非内存池需要改变），成本比调度OS线程低很多
另一方面充分利用了多核的硬件资源，近似的把若干goroutine均分在物理线程上，再加上本身goroutine的超轻量，以上种种保证了go调度方面的性能

Go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个OS线程来同时执行Go代码
默认值是机器上的CPU核心数。例如在一个8核心的机器上，调度器会把Go代码同时调度到8个OS线程上（GOMAXPROCS是m:n调度中的n）

Go语言中可以通过runtime.GOMAXPROCS()函数设置当前程序并发时占用的CPU逻辑核心数。
Go1.5版本之前，默认使用的是单核心执行。Go1.5版本之后，默认使用全部的CPU逻辑核心数。

Go语言中的操作系统线程和goroutine的关系：
一个操作系统线程对应用户态多个goroutine
go程序可以同时使用多个操作系统线程
goroutine和OS线程是多对多的关系，即m:n

*/

func a() {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}
func c() {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		fmt.Println("C:", i)
	}
}

func d() {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		fmt.Println("D:", i)
	}
}

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	runtime.GOMAXPROCS(1) // 两个任务只有一个逻辑核心，此时是做完一个任务再做另一个任务。
	go a()
	go b()
	
	// runtime.GOMAXPROCS(2) // 将逻辑核心数设为2，此时两个任务并行执行
	// go c()
	// go d()
	wg.Wait()
}
