package main

import "fmt"

// 单项通道，多用在函数的参数里
// 将通道作为参数在多个任务函数间传递，很多时候在不同的任务函数中使用通道都会对其进行限制
// 比如限制通道在函数中只能发送，或只能接收
// chan<- int是一个只写单项通道（只能对其写入int类型值），可以对其执行发送操作，而不能执行接收操作
// <-chan int是一个只读单项通道（只能从其读取int类型值），可以对其执行接收操作，而不能执行发送操作
// 函数传参及任何赋值操作中

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func square(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go square(ch2, ch1)
	printer(ch2)

}
// 关闭一个未初始化的通道会panic
// 关闭一个已经关闭的通道会panic
// 空的通道进行接收操作，会阻塞，直到执行一个发送操作，成为非空的通道，才会再执行之前的接收操作。空的不能读
// 满的通道进行发送操作，会阻塞，直到执行下一个接收操作，空出缓冲区的空间，才可以执行发送操作。满的不能写
// 通道关闭后，执行接收操作，读取数据，直到通道为空，还可以继续接收，但是只能读取该通道元素类型的零值
