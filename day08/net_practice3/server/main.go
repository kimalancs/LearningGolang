package main

import (
	"fmt"
	"net"
	"strings"
)

//Go语言实现UDP通信

// UDP协议（User Datagram Protocol）用户数据报协议
// OSI（Open System Interconnection，开放式系统互联）参考模型中一种无连接的传输层协议
// 不需要建立连接就能直接进行数据发送和接收，属于不可靠的、没有时序的通信
// 但是UDP协议的实时性比较好，通常用于视频直播相关领域。

// 使用Go语言的net包实现的UDP服务端代码

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for { // 不需要建立连接，直接收发数据
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:]) // 接收数据
		if err != nil {
			fmt.Println("read udp failed, err:", err)
			continue
		}
		fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)
		reply := strings.ToUpper(string(data[:n]))
		_, err = listen.WriteToUDP([]byte(reply), addr) // 发送数据
		if err != nil {
			fmt.Println("write to udp failed, err:", err)
			continue
		}
	}
}
