package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

/*
日志库作业
使用接口的方式实现一个既可以往终端写日志也可以往文件写日志的简易日志库


文件操作
文件是存储在外部介质上的数据集合
分为文本文件和二进制文件

os.Open()打开一个文件，输入一个字符串（"path/filename" 文件路径），返回一个*File和一个err
对得到的文件实例调用Close()方法可以关闭文件
只读

file.Read() 读取文件的方法
func (f *File) Read(b []byte) (n int, err error)
接收一个字节切片，返回读取的字节数和可能的具体错误，遇到文件末尾时会返回0和io.EOF

*/

var x int8 = 10 // 最先执行全局声明

func init() { // 再执行init函数
	fmt.Println(x)
}

func main() { // 最后再执行main函数
	fmt.Println("hello")

	// 只读方式打开当前目录下的main.go文件
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	// 关闭文件，为了防止忘记关闭，通常使用defer注册文件关闭语句
	defer file.Close()
	// 读取文件
	// var tmp = make([]byte, 128) // 指定读的长度
	// n, err := file.Read(tmp)
	// if err == io.EOF {
	// 	fmt.Println("文件读完了")
	// 	return
	// }
	// if err != nil {
	// 	fmt.Println("read file failed, err:", err)
	// 	return
	// }
	// fmt.Printf("读取了%d字节数据\n", n)
	// fmt.Println(string(tmp[:n]))

	// 循环读取，for循环读取文件中所有的数据
	// var content []byte
	// var tmp2 = make([]byte, 128)
	// for {
	// 	n, err := file.Read(tmp2)
	// 	if err == io.EOF {
	// 		fmt.Println("文件读完了")
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println("read file failed, err:", err)
	// 		return
	// 	}
	// 	content = append(content, tmp2[:n]...)
	// }
	// fmt.Println(string(content))

	// bufio 在file的基础上封装了一层API，支持更多的功能
	// bufio 按行读取
	// reader := bufio.NewReader(file)
	// for {
	// 	line, err := reader.ReadString('\n') // 注意此处是字符
	// 	if err == io.EOF {
	// 		if len(line) != 0 {
	// 			fmt.Println(line)
	// 		}
	// 		fmt.Println("文件读完了")
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println("read file failed, err:", err)
	// 		return
	// 	}
	// 	fmt.Print(line) // 按行输出，不需要Println
	// }

	// ioutil 读取整个文件
	// ioutil.ReadFile方法可以读取完整的文件

	content2, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content2))
}
