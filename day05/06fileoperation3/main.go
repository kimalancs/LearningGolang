package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
获取用户输入时，输入内容有空格
fmt.Scanln()检测到空格就结束输入
使用bufio解决
*/
func useScan() {
	var s string
	fmt.Print("请输入内容：")
	fmt.Scanln(&s)
	fmt.Printf("你输入的内容是：%s\n", s)
}

func useBufio() {
	var s string
	fmt.Print("请输入内容：")
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	fmt.Printf("你输入的内容是：%s\n", s)
}

func main() {
	// useScan() // 输入的内容只有第一个空格前的内容被Scanln()函数获取
	useBufio() // 输入的内容可以有空格
}
