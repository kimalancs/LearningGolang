package main

import (
	"fmt"
	"path"
	"runtime"
)

// 0 返回的就是Caller代码所在的那一行的行数、当前的函数名、当前文件的文件名
// 1 返回的就是Caller代码的上一层的调用函数的行数，上一层的函数名、上一层的文件名
// 2 返回的就是Caller代码的上两层的

func f() {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName) // 0,返回的是main.f, 1,返回的是main.main, 2，返回的是runtime.main
	fmt.Println(file)
	fmt.Println(line) // 当runtime.Caller()括号里是1时，line返回的就是Caller代码的上一层，调用f()函数的那一行的行数

}

func f1() {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	fmt.Println(file)
	fmt.Println(line) // 当runtime.Caller()括号里是2时，line返回的就是Caller代码的上两层，调用f2()函数的那一行的行数
}

func f2() {
	f1()
}

func main() {
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	fmt.Println(file)
	fmt.Println(path.Base(file)) // 拿到路径最后的文件名
	fmt.Println(line) // 当runtime.Caller()括号里是0时，line返回的就是现在调用Caller的那一行的行数

	f()
	f1()
	f2()

}
