package main

import "fmt"

// struct 结构体
// 值类型
// 没有类的概念，也不支持类的继承等面向对象的概念
// 通过结构体的内嵌再配合接口比面向对象具有更高的扩展性和灵活性
// 基础数据结构一般都是单一维度的基本属性，想要表达一个事物的多个属性时，就可以自定义结构体，封装多个基本数据类型
// 内置基本数据类型是用来描述一个值的，而结构体是用来描述一组值的，本质上是一种聚合型的数据类型

type person struct {
	name string // 类型相同的字段可以写在一行. name, city string
	city string
	age  int8
}


func main() {
	// 结构体实例化后才会分配内存，才能使用其中的字段
	var p1 person // var 结构体实例 结构体类型
	p1.name = "kim" // 通过.来访问结构体的字段
	p1.city = "Beijing"
	p1.age = 19
	fmt.Printf("p1=%v\n", p1)

	// 需要定义一些临时数据结构的场景下，可以使用匿名结构体
	var user struct{Name string; Age int} // 结构体字段可以写在一行，用分号隔开
	user.Name = "Taeyeon"
	user.Age = 18
	fmt.Printf("%#v\n", user)

	var num1 = struct {x int;y int} {10, 20} // 定义并初始化
	fmt.Println(num1)

	// 结构体指针，使用new关键字对结构体进行实例化，得到的是结构体的地址
	var p2 = new(person)
	fmt.Printf("%T\n", p2)
	fmt.Printf("p2=%#v\n", p2)
	fmt.Printf("%v\n", p2)
	// 支持对结构体指针直接使用.来访问结构体的成员
	// p2.age = 28 在底层是(*p2).name =28，这是Go语言帮我们实现的语法糖
	p2.name = "allen"
	p2.age = 28
	// (*p2).age = 19
	p2.city = "Seoul"
	fmt.Printf("%#v\n", p2)
	fmt.Printf("%v\n", p2)
	fmt.Printf("%p\n", p2)
	fmt.Printf("%p\n", &p2)
	// 使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作
	// 没有初始化的结构体，其成员变量都是对应其类型的零值
	p3 := &person{}
	fmt.Printf("%T\n", p3)
	fmt.Printf("%#v\n", p3)
	// 结构体传入函数，跟值类型一样，都是复制，函数内改的是副本，不影响原变量，需要改原变量就要传入指针（内存地址）
	changeAge := func(x *person){
		x.age = 18
	}
	changeAge(p2)
	fmt.Println(p2.age) // 18， 传入指针可以修改原变量
	p4 := person{ //使用键值对的时候不用管顺序
		name: "tae",
		age: 17,
		city: "shanghai",
	}
	fmt.Println(p4)
	p5 := person{ //值列表方式，不写key只写value时，要按结构体定义的字段顺序，且必须初始化所有字段，不能与键值对方法混用
		"yoona",
		"shenzhen",
		19,
	}
	fmt.Println(p5)
	p6 := person { // 可以只初始化部分字段，没有指定的字段为该字段类型的零值
		city: "Taipei",
	}
	fmt.Println(p6)
	// 结构体占用一块连续的内存
	var mem struct {
		a int8
		b int8
		c int8
		d int8
	}
	mem.a = 1
	mem.b = 2
	mem.c = 3
	mem.d = 4
	fmt.Printf("mem.a %p\n", &mem.a)
	fmt.Printf("mem.b %p\n", &mem.b)
	fmt.Printf("mem.c %p\n", &mem.c)
	fmt.Printf("mem.d %p\n", &mem.d)

	// 构造函数
	// 返回一个结构体变量的函数
	// Go语言没有内置的构造函数，需要自己实现
}
