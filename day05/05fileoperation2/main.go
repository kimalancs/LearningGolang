package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/*
文件写入
os.OpenFile()函数能以指定模式打开文件，从而实现文件写入相关功能
func OpenFile(name string, flag init, perm FileMode) (*File, error)
name 要打开的文件名，路径
flag 打开文件的模式
perm 文件权限，一个八进制数，读 r ==> 04, 写 w ==> 02, 执行 x ==> 01

flag
os.O_WRONLY	只写
os.O_CREATE	创建
os.O_RDONLY	只读
os.O_RDWR	读写
os.O_TRUNC	清空
os.O_APPEND	追加

借助io.Copy()实现一个拷贝文件函数

*/

// CopyFile copy a file
func CopyFile(dstName, srcName string) (written int64, err error) {
	// 只读方式打开源文件
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("open %s failed, err: %s", srcName, err)
		return
	}
	defer src.Close()
	// 以写｜创建的方式打开目标文件
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open %s failed, err: %s", dstName, err)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src) // 调用io.Copy()拷贝内容
}

func main() {
	file, err := os.OpenFile("./xxx.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	str := "hello\n"
	file.Write([]byte(str)) // 写入字节切片数据

	file.WriteString("你好\n") // 直接写入字符串数据

	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello\n") // 将数据写入内存
	}
	writer.Flush() // 将缓存中的内容写入文件

	str2 := "taeyeon\n"
	err2 := ioutil.WriteFile("./xxx2.txt", []byte(str2), 0666)
	if err2 != nil {
		fmt.Println("write file failed, err:", err)
		return
	}

	_, err3 := CopyFile("./xxx3.txt", "./xxx2.txt")
	if err3 != nil {
		fmt.Println("copy file failed, err:", err)
		return
	}
	fmt.Println()

}
