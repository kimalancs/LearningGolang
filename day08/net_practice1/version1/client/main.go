package main

import (
	"fmt"
	"net"
)

// tcp client

func main() {
	// 与server建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("dial 127.0.0.1:20000 failed", err)
		return
	}

	// 发送数据
	conn.Write([]byte("hello"))
	conn.Close()
}
