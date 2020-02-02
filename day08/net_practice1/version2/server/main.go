package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp server version2

func processConn(conn net.Conn) {
	defer conn.Close()
	// 与客户端通信
	var tmp [128]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("read from client failed", err)
			return
		}
		fmt.Println(string(tmp[:n]))
		fmt.Println("请回复：")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}
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
		go processConn(conn)
	}
}
