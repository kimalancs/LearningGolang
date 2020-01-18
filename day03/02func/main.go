package main

import (
	"fmt"
	"strings"
)

// 闭包指的是一个函数和与其相关的引用环境组合而成的实体
// 闭包=函数+引用环境
// 这个引用环境相当于闭包函数的预设条件，修改这个预设条件，可以生成功能不同的闭包函数
// 闭包函数是函数的返回值，一个匿名函数
// 预设条件是函数的参数，通过修改参数，改变预设条件，返回不同的匿名函数，最终生成多种功能的闭包函数

//例1 adder函数没有参数，预设条件x初始化时为0，之后每次运行闭包函数，都会改变预设条件
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

//例2 需要把f2作为参数传入f1执行，由于f2有两个参数，f1的参数是没有参数的函数（类型不匹配）,不能直接传。
// 定义一个函数f3包装f2，然后传入f1。
func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

func f3(f func(int, int), x, y int) func() {
	tmp := func() {
		f(x, y)
	}
	return tmp
}

// 例3 adder2函数的参数x为预设条件
func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

// 例4 makeSuffixFunc函数的suffix参数是预设条件，通过改变suffix，生成添加.jpg,.txt等多种功能的闭包函数
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 例5，返回的两个闭包函数，add、sub，每次运行都会修改公用的base，下次运行就在上次修改的基础上进行操作
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	var f = adder()
	fmt.Println(f(10)) // 10
	fmt.Println(f(20)) // 30
	fmt.Println(f(30)) // 60
	// 变量f是一个函数并且引用了其外部作用域中的x变量，此时f就是一个闭包
	// 在f的生命周期内，变量x也一直有效

	ret := f3(f2, 100, 200)
	f1(ret)

	f1 := adder2(200)
	fmt.Println(f1(10))
	fmt.Println(f1(20))

	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test"))
	fmt.Println(txtFunc("test"))
	fmt.Println(txtFunc("test.jpg"))

	c1, c2 := calc(10)
	fmt.Println(c1(1), c2(2)) // 11 9
	fmt.Println(c1(3), c2(4)) // 12 8
	fmt.Println(c1(5), c2(6)) // 13 7
	// base每次调用闭包函数都会修改

}
