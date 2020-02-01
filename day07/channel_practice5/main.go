package main

import "fmt"

/* select多路复用
同时从多个通道接收数据
通道在接收数据时，如果没有数据可以接收，会发送阻塞
通过遍历的方式可以实现，但是运行性能会差很多
for {
	// 尝试从ch1接收值，先从ch1接收值，直到取完，然后if判断再从ch2接收，不是随机的
	data, ok := <- ch1
	// 尝试ch2接收值
	data,ok := <- ch2
	...
}

为了应付这种场景，内置select关键字，可以同时响应多个通道的操作
select类似switch语句，有一系列case分支和一个默认的分支
每个case会对应一个通道的通信（接收或发送）过程
select会一直等待，直到某个case的通信操作完成时，就会执行case分支对应的语句

select{
case <-ch1:
	...
case data := <- ch2:
	...
case ch3<-data:
	...
default:
	default
}

select可以提高代码的可读性

可处理一个或多个channel的发送/接收操作
如果多个case同时满足，select会随机选择一个
对于没有case的select{}会一直等待，可用于阻塞main函数


*/
func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}

}
