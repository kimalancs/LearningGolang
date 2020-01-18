package main

import "fmt"

//defer语句会将其后跟随的语句进行延迟处理，在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行
//也就是说先被defer的语句最后被执行，最后被defer的语句最先被执行，后进先出

// return语句在底层并不是原子操作，分为两步，返回值赋值、RET指令
// defer语句执行的实际就在返回值赋值操作后，RET指令执行前

// defer多用于函数结束之前释放资源，socket连接，数据库连接，资源清理，记录时间，解锁等

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x // 5 return返回值赋值5，之后执行defer改变了x，而返回值不变
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 // 6 return赋值返回值x=5，之后执行defer改变了x，此时x是返回值，所以返回值也改变了
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x // 5 return赋值返回值y=5，之后执行defer改变了x，此时y是返回值，不变
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 // 5 return赋值返回值x=5，之后执行defer，将x作为参数传入defer语句中的函数，此时改变的是x的副本，而返回值x不变
}

func f5() (x int) {
	defer func(x *int) {
		*x++
	}(&x)
	return 5 // 6 return返回值赋值x=5，之后执行defer，传入指针时，修改的是变量x的原本，返回值x变为6
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
	fmt.Println(f5())

	fmt.Println("start")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("End")

	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y)) //defer注册时，第二个参数x确定为1，为确定第三个参数，先执行内部calc函数，输出A 1 2 3
	x = 10
	defer calc("BB", x, calc("B", x, y)) //defer注册时，第二个参数确定为10，为确定第三个参数，先执行内部calc函数，输出B 10 2 12
	y = 20 // defer之后变量的变化已经不影响defer最后执行的结果了
	// 提示：defer注册要延迟执行的函数时该函数所有的参数都要确定其值

}
