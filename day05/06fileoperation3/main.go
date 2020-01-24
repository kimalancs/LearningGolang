package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
练习1
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
	// fmt.Println("请输入内容：")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入内容：")
	s, _ = reader.ReadString('\n')
	fmt.Printf("你输入的内容是：%s\n", s)
}

/*
练习2
在文件中间插入内容
没办法直接插入，插入的内容会覆盖原本位置的内容

插入内容之前的部分为A
插入内容之后的部分为B
在AB之间插入内容C

先建一个临时文件，写入A，再写入C，再写入B，最后删除源文件，将临时文件改为源文件的名

*/
func insertString() {
	fileObj, err := os.OpenFile("./xxx.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open file failed, err: %v", err)
		return
	}

	tmpFile, err := os.OpenFile("./xxx.tmp", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("create file failed, err: %v", err)
		return
	}
	defer tmpFile.Close()

	var ret [6]byte
	n, err2 := fileObj.Read(ret[:])
	if err2 != nil {
		fmt.Printf("read from file failed, err: %v", err2)
		return
	}

	tmpFile.Write(ret[:n])

	str := "你好\n"
	tmpFile.Write([]byte(str))

	var x [128]byte
	for {
		n, err := fileObj.Read(x[:])
		if err == io.EOF {
			tmpFile.Write(x[:n])
			break
		}
		if err != nil {
			fmt.Printf("file read failed, err: %v", err)
			return
		}
		tmpFile.Write(x[:n])
	}
	fileObj.Close()
	tmpFile.Close()
	os.Rename("./xxx.tmp", "./xxx.txt")
}

func main() {
	// useScan() // 输入的内容只有第一个空格前的内容被Scanln()函数获取
	// useBufio() // 输入的内容可以有空格

	insertString()

}
