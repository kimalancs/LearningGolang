package main

import "fmt"

// type 关键字
//自定义类型，始终生效
type myInt int

//类型别名，只是代码编写时生效，代码编译后就失效，编写代码时更清晰
type yourInt = int

func main() {
	var n1 myInt
	n1 = 10
	fmt.Printf("%v %T\n", n1, n1)

	var n2 yourInt
	n2 = 200
	fmt.Printf("%v %T\n", n2, n2)

	var m rune // rune就是int32的类型别名
	m = '中'
	fmt.Printf("%c %v %T\n", m, m, m)



}
