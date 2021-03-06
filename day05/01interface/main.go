package main

import "fmt"

/*
接口interface
定义一个对象的行为规范，只定义规范，不实现
接口是一种类型，一种抽象的类型
一组method的集合
duck-type programming的一种体现
接口指定了一个方法集，实现模块化的强大方式
方法集的蓝本，描述了方法集中的所有方法，但没有实现他们
可以在符合接口要求的前提下随便更换实现
支持多种实现，不同类型都实现了一个接口，但是实现方式可以不同

接口指定了每个方法的函数签名（方法名、参数、返回值）
实现一个接口，只要实现了指定的方法集，并且函数签名一样就行

不同类型都实现了一个接口，函数接收接口类型参数，可以编写可重用于多个类型的函数，有助于代码重用

如编写一个使用MySQL数据库的程序，不使用接口时，代码完全针对MySQL，如果要换成其他数据库需要重写大量代码
而定义一个数据库接口，只要满足接口的要求，就可以使用任何数据库
接口可包含多个实现，引入了多态的概念
一个方法集实现了一个接口，就可以说它与另一个实现该接口的方法集互为多态


不关心一个变量是什么类型，只关心这个变量能调用的方法
一台机器，只要有洗衣服的功能，就叫洗衣机
汽车可以跑，动物可以跑，只要会跑，就可以写成Mover接口，多个类型都可以实现Mover接口的方法，都属于Mover接口类型
不关心属性（数据），只关心行为（方法）
减少重复的代码，具备相同行为的类型，要为每个类型都写一个该行为的方法，就会产生很多冗余代码，使用接口可以给多个类型定义一个行为的方法

type 接口类型名 interface{
	方法1（参数列表1）（返回值列表1）
	方法2（参数列表2）（返回值列表2）
	...
}

接口命名一般在最后加er，接口名最好能体现该接口的类型含义
有写操作的接口叫Writer
有字符串功能的接口叫Stringer

参数列表和返回值列表都可以省略

方法名和接口类型名都首字母都是大写时，这个方法可以被接口所在包之外的代码访问

接口中的方法必须全部实现的类型才能算实现了这个接口，才能算是这个接口类型

一个对象只要实现了接口中的所有方法，那么就实现了这个接口

接口是需要实现的方法列表

接口类型变量能存储所有实现了该接口的实例

像类型一样使用即可，给变量、参数、返回值指定类型，比如函数参数和返回值设置成接口类型，实现了接口的类型都可以传入或返回

值接收者和指针接收者实现接口有区别
值接收者实现接口，值类型还是指针类型变量都可以赋值给接口变量
指针接收者实现接口，只有指针类型变量可以赋值给接口变量，值类型变量不可以赋值给接口变量，相当于实现该接口的是指针类型，值类型没有实现

如果实现接口的是指针接收者，那么接收该接口类型的函数，只能接收该类型变量的指针，而不能接收该类型变量


*/
type barker interface {
	bark()
}

type dog struct{}

type cat struct{}

type human struct{}

func (c cat) bark() { // 值接收者实现接口，不管是cat类型还是*cat类型的变量，都可以赋值给该接口变量
	fmt.Println("喵喵喵")
}

func (d dog) bark() {
	fmt.Println("汪汪汪")
}

func (h *human) bark() { // 指针接收者实现接口，*human类型的变量可以赋值给接口变量，human类型不可以
	fmt.Println("Help")
}

func alarm(x barker){ // 只要实现了bark方法的类型都可以barker接口类型的参数传入alarm函数
	x.bark()
}

func main() {
	var c cat
	var d = &dog{} // 变量d是dog类型结构体指针
	var x barker // 接口类型变量
	
	alarm(c)
	alarm(d)

	x = c // 只要实现了该接口所有方法的实例都可以存储在该接口类型变量中
	alarm(x)

	x = d // 结构体指针也可以赋值给接口变量
	alarm(x)

	var h1 = human{}
	// x = h1 //报错
	fmt.Printf("%T\n", h1)
	// alarm(h1) // 报错
	// 实现barker接口的是*human类型，而不是human类型

	var h2 = &human{}
	x = h2
	alarm(h2)


}
