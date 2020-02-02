package main

import (
	"fmt"
	"net"
)

// 网络编程
// socket编程
// socket，套接字，UNIX进程通信机制，用于描述IP地址和端口，是一个通信链的句柄
// Socket可以理解为TCP/IP网络的API，它定义了许多函数或例程，程序员可以用它们来开发TCP/IP网络上的应用程序
// 电脑上运行的应用程序通常通过”套接字”向网络发出请求或者应答网络请求
// Socket是应用层与TCP/IP协议族通信的中间软件抽象层
// 在设计模式中，Socket其实就是一个门面模式
// 它把复杂的TCP/IP协议族隐藏在Socket后面
// 对用户来说只需要调用Socket规定的相关函数，让Socket去组织符合指定的协议数据然后进行通信

// 个TCP服务端可以同时连接很多个客户端，例如世界各地的用户使用自己电脑上的浏览器访问淘宝网
// 因为Go语言中创建多个goroutine实现并发非常方便和高效，所以我们可以每建立一次链接就创建一个goroutine去处理

// TCP服务端程序的处理流程：
// 监听端口
// 接收客户端请求建立链接
// 创建goroutine处理链接


// tcp server

func processConn(conn net.Conn){
	// 与客户端通信
	var tmp [128]byte
	n, err := conn.Read(tmp[:])
	if err != nil {
		fmt.Println("read from conn failed", err)
		return
	}
	fmt.Println(string(tmp[:n]))
}

func main() {
	// 本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("start server on 127.0.0.1:20000 failed", err)
		return
	}
	// 等待别人来跟我建立连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed", err)
			return
		}
		go processConn()
	}
}
