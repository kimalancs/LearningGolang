package main

import (
	"errors"
	"fmt"
)

/*
	函数是组织好的、可重复使用的、用于执行指定任务的代码库
	一段代码的封装
	把一段逻辑抽象出来封装到一个函数中，起一个名字，每次用到就直接用函数名调用即可
	使用函数让代码结构更清晰、更简洁

	func 函数名(参数)(返回值){
		函数体
	}

	参数由参数变量和参数变量的类型组成
	多个参数之间用,英文逗号分隔

	返回值由返回值变量和变量类型组成，也可以只写返回值的类型
	多个返回值必须用()包裹，并用,英文逗号分隔
	只有一个返回值时只写返回值类型则可不用()

	函数体是实现指定功能的代码块


*/

func sum(x int, y int) (ret int) { // 命名返回值相当于在函数中声明了一个变量，可以直接在函数体中使用该变量
	ret = x + y
	return
}

func sum2(x int, y int) int { //返回值可以命名也可以不命名
	return x + y
}

func f1(x, y, z int, m, n string) { // 多个参数类型相同时可以非最后一个参数的类型省略。类型简写，相邻变量类型相同时，可省略类型
	fmt.Println(x + y + z)
	fmt.Println(m + n)
}

func f2(x int, y int) { // 没有返回值
	fmt.Println(x + y)
}

func f3() { // 没有参数也没有返回值
	fmt.Println("f2")
}

func f4() int { // 没有参数，有返回值
	return 3
}

func f5() (int, string) { // 可以返回多个值
	return 309, "Taeyeon"
}

func f6(x string, y ...int) { // 可变长参数，参数数量不固定，必须放在函数参数的最后，固定参数放在前边，通过切片实现
	fmt.Println(x)
	fmt.Println(y) // y的类型是slice
}

//函数没有默认参数的概念，提倡简洁，要明确的显式的传递，不要晦涩的

// 调用有返回值的函数时，可以不接收其返回值

// 变量作用域
// 全局变量，定义在函数外部，程序整个运行周期内都有效，函数中可以访问到全局变量
//局部变量分两种，函数内定义的变量和语句块定义的变量
// 函数内定义的变量无法在该函数外使用，仅在该函数内生效
// 函数内定义的局部变量与全局变量重名，该函数中优先使用局部变量
// 函数中查找变量的顺序，先查内部的局部变量，再找外层的局部变量，直到全局变量，再找不到就报错
// 语句块定义的变量（if条件判断、for循环、switch语句中）只在语句块中生效

//type关键字来定义一个函数类型
type calculation func(int, int) int //此处定义了一种函数类型，这种函数接收两个int类型的参数，并返回一个int类型的返回值

func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}

func calc(x, y int, op func(int, int) int) int { ///函数可以作为参数
	return op(x, y)
}

func do(s string) (func(int, int) int, error) { // 函数也可以作为返回值
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

func main() {
	var c calculation // 可以声明函数类型的变量
	c = add           // 可将add函数赋值给calculation类型的变量
	fmt.Printf("%T\n", c)

	ret := calc(10, 20, add)
	fmt.Println(ret) // 30

	// 匿名函数
	// 函数内部不能定义有名字的函数，只能定义匿名函数（没有函数名的函数）
	// func(参数)(返回值){
	//	函数体
	// }
	// 由于没有函数名，不能像普通函数一样调用，需要保存到某个变量，或者作为立即执行函数
	// 多用于实现回调函数和闭包
	add2 := func(x, y float64) { //将匿名函数保存到变量
		fmt.Println(x + y)
	}
	add2(1.5, 2.5) // 通过变量调用匿名函数

	//如果只执行一次的函数，还可以简化成立即执行函数，函数定义完，加()直接执行，括号内可加参数
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)

}
