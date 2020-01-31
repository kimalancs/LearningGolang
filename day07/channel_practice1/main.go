package main

import (
	"fmt"
	"sync"
)

// 单纯函数并发执行没有意义，函数与函数间需要交换数据才能体现并发执行函数的意义
// 有些语言使用共享内存来实现通信，多个进程访问同一内存时，需要通过排他锁来保证内存安全，加锁解锁会造成性能问题
// 通过通信来共享内存，channel是goroutine之间连接，channel让一个goroutine发送特定值到另一个goroutine
// channel通道是一种特殊类型，像一个单向传送带或队列，总是遵循先入先出first in first out的规则，保证收发数据的顺序
// 每一个通道都是一个具体类型的导管，就是声明channel时需要为其指定元素类型
// var 变量 chan 元素类型
// var ch1 chan int //声明一个传递整形的通道
// 通道是引用类型，空值是nil
// 声明通道之后，需要make函数初始化才能使用
// make(chan 元素类型, [缓冲大小])

// 通道有发送（send）、接收(receive）和关闭（close）三种操作。
// 发送和接收都使用<-符号

var ch1 chan int

var wg sync.WaitGroup

func main() {
	fmt.Println(ch1)
	ch1 := make(chan int) //
	fmt.Println(ch1)
	// ch1 <- 1 // 可以通过编译，但是执行时会报错，fatal error: all goroutines are asleep - deadlock!
	// 无缓冲的通道又称为阻塞的通道
	// 只有在有goroutine接收值时才能发送值

	// 一种方法是，启用一个goroutine来接收值，如下
	wg.Add(1)
	go func() { //此goroutine会在后台等待，等到下一步ch1接收到值后再执行
		defer wg.Done()
		x := <-ch1 // 从ch1中接收值，并赋值给x
		fmt.Println("后台goroutine从通道ch1中接收了值", x)
	}()
	ch1 <- 10 // 把10发送给通道ch1
	fmt.Println("10发送到通道ch1中了")
	// 无缓冲通道上的发送操作会阻塞，直到另一个goroutine在该通道上执行接收操作，这时值才能发送成功，两个goroutine将继续执行
	// main函数也是一个goroutine
	// 如果接收操作先执行，接收方的goroutine会阻塞，直到另一个goroutine在该通道上发送一个值
	// 使用无缓冲通道进行通信将导致发送和接收的goroutine同步化，因此无缓冲通道也被称为同步通道

	close(ch1) // 关闭通道ch1
	// 只有在通知接收方goroutine所有的数据都发送完毕的时候才需要关闭通道
	// 通道是可以被垃圾回收机制回收的，它和关闭文件是不一样的
	// 在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的。
	// 关闭后的通道有以下特点：
	// 对一个关闭的通道再发送值就会导致panic。
	// 对一个关闭的通道进行接收会一直获取值直到通道为空
	// 对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值
	// 关闭一个已经关闭的通道会导致panic
	wg.Wait()

	//另一种方法是使用有缓冲区的通道，在make函数初始化通道时，指定通道的容量
	var ch2 = make(chan int, 2) // 创建容量为2的有缓冲区的通道
	fmt.Println(ch2)
	fmt.Println(&ch2)
	ch2 <- 11 // 有缓冲区的通道可以接收值，但超出缓冲区大小之后，还会发生deadlock
	// 只要通道的容量大于零，那么该通道就是有缓冲的通道
	// 通道的容量表示通道中能存放元素的数量
	// 可以使用内置len()函数获取通道内元素的数量，cap()函数获取通道的容量，但很少这么做
	x := <-ch2 // x取到的是ch2中的int值，所以x是int类型
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	// <-ch1      // 从ch1中接收值，忽略结果

	// 当通道被关闭时，再往该通道发送值会引发panic，从该通道里接收的值一直都是类型零值。那如何判断一个通道是否被关闭了，如下两种方法
	ch3 := make(chan int)
	ch4 := make(chan int)
	// 开启一个goroutine，将0-99发送到ch3中
	go func() {
		for i :=0; i <100; i++ {
			ch3 <- i
		}
		close(ch3)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch3 // 接收时会返回一个值和一个布尔值，布尔值可以接收也可以不接收
			// 通道内元素取完之前，返回true，取完后后，再取就返回false
			if !ok {
				fmt.Println(i) // 元素取完后，再取值返回对应类型的零值
				break
			}
			ch4 <- i * i
		}
		close(ch4)
	}()
	// 在主goroutine中从ch4中接收值打印
	for i := range ch4 { // 通道关闭后会退出for range循环
		fmt.Println(i)
	}
	// 通常使用的是for range从通道中循环取值
	// 使用for range遍历通道，当通道被关闭的时候就会退出for range

	ch5 := make(chan string, 2)
	ch5 <- "hello"
	ch5 <- "world"
	close(ch5)
	for {
		i, ok := <-ch5
		if !ok {
			fmt.Printf("%#v\n", i) // 取完后
			break
		}
	}
}
