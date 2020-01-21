package main

import "fmt"

/*
多个类型可以实现同一个接口
每个类型都要实现这个接口的所有方法，可以多不可以少

一个类型可以实现多个接口
接口之间彼此独立，不知道对方的实现

一个接口的方法不一定需要由一个类型完全实现，可以通过在类型中嵌套其他类型或者结构体来实现
如athlete类型实现了mover接口的方法，将他嵌入到其他类型，其他类型就自动实现了mover接口
或者某个接口定义了三个方法，A类型实现了两个方法，B类型实现了另一个方法，将A嵌入B，则B就实现了这个接口

接口嵌套
接口里包含接口，比如mover接口，包含runer接口，walker接口，swimer接口
嵌套的接口使用和普通接口一样，比如mover接口类型的变量mover1.run()直接调用runer接口中的run方法

空接口，没有定义任何方法的接口，任何类型都实现了空接口，所以空接口类型的变量可以存储任何类型的变量
interface{}
函数参数类型写成空接口，就可以接收所有类型的参数

一个接口的值由一个具体类型和具体类型的值两部分组成
动态类型
动态值
声明一个接口变量，此时两部分都是nil
赋值给这个变量，此时才确定类型和值，这样才能存储不同类型的变量

接口断言
判断空接口中的值是什么类型
x.(T)
x表示类型为interface{}的变量
T表示断言x可能是的类型
为什么要有接口断言：传入参数是空接口类型，想要判断到底是什么类型

只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口
不要为了接口而写接口，这样只会增加不必要的抽象，导致不必要的运行时损耗

*/
//断言多次时可以使用switch语句
func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string, value is %v\n", v)
	case int:
		fmt.Printf("x is a int, value is %v\n", v)
	case bool:
		fmt.Printf("x is a bool, value is %v\n", v)
	default:
		fmt.Println("unsupport type")
	}
}

type runer interface {
	run()
}

type walker interface {
	walk()
}

type swimer interface {
	swim()
}

type mover interface {
	runer
	walker
	swimer
}

type athlete struct {
	name string
}

type proath struct {
	athlete
}

func (a athlete) swim() {
	fmt.Printf("%v can swim\n", a.name)
}

func (a athlete) run() {
	fmt.Printf("%v can run\n", a.name)
}

func (a athlete) walk() {
	fmt.Printf("%v can walk\n", a.name)
}

// Peopler people
type Peopler interface {
	Speak(string) string
}

// Student student
type Student struct{}

// Speak speak
func (s *Student) Speak(think string) (talk string) {
	if think == "idiot" {
		talk = "You're so good at it"
	} else {
		talk = "hello"
	}
	return
}

func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func main() {
	var peo Peopler
	fmt.Printf("%T\n", peo) // <nil>, 此时peo没有明确的类型
	peo = &Student{}
	fmt.Printf("%T\n", peo) // *main.Student, 此时peo有了明确的类型
	think := "bitch"
	fmt.Println(peo.Speak(think))
	// 接口类型的值分成类型和值两个部分，动态类型、动态值
	// 一开始两部分都是nil，当赋值后，才确定类型和值，这样才能存储不同类型的变量

	var ath mover
	ath = athlete{}
	ath.walk()
	ath.run()
	ath.swim()

	var proath1 mover
	proath1 = proath{athlete: athlete{name: "liuxiang"}}
	proath1.walk()
	fmt.Printf("%T\n", proath1)
	fmt.Printf("%#v\n", proath1)

	show(proath1)

	var x interface{}
	x = 18
	v, ok := x.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}

	justifyType(18)
}
