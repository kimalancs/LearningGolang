package main

import (
	"fmt"
	"time"
)

// Go天生支持并发，通过goroutine实现
// 并发concurrency：同一时间段内执行多个任务，多个任务切换执行
// 并行parallelism：同一个时刻执行多个任务，需要硬件支持，多核CPU才能真正实现
// concurrency is about dealing with lots of things at once
// parallelism is about doing lots of things at once
// not the same, but related
// concurrency is about structure, parallelism is about execution
// concurrency provides a way to structure a solution to solve a problem that may (but not necessarily) be parallelizable
// 并发不同于并行，但并发也许可以让程序实现并行化

// 如果程序属于CUP密集型，使用并发，无法提升程序性能，反而因为大量计算资源花在创建线程本身，而导致程序性能进一步下降
// 如果程序属于IO密集型，在进行程序压测的时候发现CPU占用率很低， 但性能却遇到了瓶颈，原因是程序将大量的时间花在了等待IO的过程中
// 在等待IO的时候继续执行其他的程序逻辑即可提高CPU利用率，从而提高我们的程序性能，此时并发的好处就体现出来了

// goroutine 类似线程，属于用户态的线程，可以根据需要创建成千上万个goroutine并发工作
// goroutine是由Go语言的运行时（runtime）调度完成，而线程由操作系统调度完成
// Java或C++实现并发时，通常要自己维护一个线程池，并且自己去包装一个又一个的任务，自己去调度线程执行任务并维护上下文切换
// goroutine由runtime调度和管理，智能地将goroutine中的任务合理分配给每个CPU，内置调度和上下文切换的机制
// GMP是Go运行时层面的实现，是Go语言自己实现的一套调度系统，区别于操作系统调度OS线程
// G，goroutine，协程，每个goroutine对象中的sched保存其上下文信息，与所在P的绑定信息
// M，machine，对内核级线程的封装，数量对应真实的CPU数（真正干活的对象），goroutine最终放到M上执行
// P，processor，G和M的调度对象，用来调度G和M之间的关联关系，其数量可以通过runtime.GOMAXPROCS()来设置，默认为核心数
// P管理一组goroutine队列，P里面会存储当前goroutine运行的上下文环境（函数指针，堆栈地址及地址边界）
// P会对自己管理的goroutine队列做一些调度（比如把占用CPU时间较长的goroutine暂停、运行后续的goroutine等等）
// 当自己的队列消费完了就去全局队列里取，如果全局队列里也消费完了会去其他P的队列里抢任务
// P与M一般也是一一对应的。他们关系是：P管理着一组G挂载在M上运行。当一个G长久阻塞在一个M上时，runtime会新建一个M，阻塞G所在的P会把其他的G挂载在新建的M上
// 当旧的G阻塞完成或者认为其已经死掉时回收旧的M。


// goroutine极其轻量级，在初始化时通常仅需要分配2KB的栈空间，而线程通常需要占用1MB
// OS线程（操作系统线程）一般都有固定的栈内存，通常为2MB，一个goroutine的栈在其生命周期开始时只有很小的栈，典型情况下2KB
// goroutine的栈不是固定的，可以按需增大和缩小，goroutine的栈大小限制可以达到1GB，虽然极少会用到这么大
// 一次创建十万左右的goroutine也是可以的


// 使用channel在多个goroutine间进行通信
// CSP并发模式 Communicating Sequential Process

// 在调用函数的时候在前面加上go关键字，就可以为一个函数创建一个goroutine
// 将一个任务包装成一个函数，然后为其创建goroutine

func hello() {
	fmt.Println("hello")
}

func main() {
	go hello()
	for i := 0; i < 100; i++ {
		go func() { // 此时是闭包，i是函数外部变量，要向外部获取，当for循环中i变化快时，goroutine内闭包函数去获取时已经变化了，可能出现多个相同的i
			fmt.Println(i)
		}()
	}

	for i := 0; i < 100; i++ {
		go func(i int) { // i作为参数传入匿名函数，不再是闭包，不会出现重复的i，但是还是乱序，goroutine调度是随机的
			fmt.Println(i)
		}(i)
	}
	fmt.Println("main")
	time.Sleep(time.Second) // 创建goroutine会花费一定时间，main函数返回时，由main函数启动的goroutine也一同结束
	// 使用time.Sleep来等1s，最粗暴的方法，时间有可能多了，也可能不足，可以使用sync.WaitGroup实现
}
