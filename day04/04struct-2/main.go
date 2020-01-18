package main

import "fmt"

// 构造函数
// 返回一个结构体变量的函数
// Go语言没有内置的构造函数，需要自己实现

// 标识符首字母大写，对外部可见，公有，能被其他包调用。首字母小写，外部不可见，私有，导入了包也不能用。
// 公开的函数需要在其上写注释，且有格式要求
type person struct {
	name string
	age  int
}

// 结构体是值类型，赋值、传参数是复制拷贝，当结构体字段很多的时候，会占用大量内存，尽量使用指针，减少开销
// 一般newPerson这种函数名就是构造函数，约定俗成
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

// 方法method是一种作用于特定类型变量的函数
// 这种特定类型变量叫做接收者receiver（类似其他语言中的this或self）
// func（接收者变量 接收者类型） 方法名（参数列表）（返回列表）{
//	函数体
//}
// 接收者变量名在命名时建议使用接收者类型名的第一个小写字母，而不是self、this之类的命名，例如Person类型的接收者变量名p
// 方法和函数的区别：函数不限定于任何类型，方法属于特定的类型
func (p person) Dream() {
	fmt.Printf("%s的梦想是唱一辈子歌\n", p.name)
}

// 指针类型的接收者由一个结构体的指针组成，由于指针的特性，调用方法时修改接收者指针的任意成员变量，在方法结束后，修改都是有效的
func (p *person) setAge(newAge int) {
	p.age = newAge
	fmt.Println(p.age)
}

// 值类型的接收者，传入方法时，拷贝接收者变量的值，修改操作只针对副本，无法修改接收者变量本身
func (p person) setAge2(newAge int) {
	p.age = newAge
	fmt.Println(p.age)
}

// 什么时候使用指针类型接收者
// 需要修改接收者中的值
// 接收者时拷贝代价比较大的大对象
// 保证一致性，如果某个方法使用了指针接收者，其他方法也应该使用指针接收者

// 接收者类型可以是任何类型
// 任何类型都可以拥有方法

// 非本地类型不能定义方法，不能给别的包的类型定义方法，可以先把该方法定义成新的自定义类型，然后再为自定义类型添加方法

// 结构体的匿名字段
// 结构体允许其成员字段在声明时没有字段名，只有类型，这种没有名字的字段就叫匿名字段
// 匿名字段默认采用类型名作为字段名，结构体要求字段名称必须唯一，所以一个机构体中同种类型的匿名字段只能有一个

type student struct {
	string
	int
}

type myInt int

func (i myInt) hello() {
	fmt.Println("我是一个int")
}

func main() {
	p1 := newPerson("kim", 18)
	p2 := newPerson("taeyeon", 17)
	fmt.Println(p1, p2)

	p2.Dream()

	p1.setAge(22)
	fmt.Println(p1.age)

	p1.setAge2(25)
	fmt.Println(p1.age)

	var s1 student
	s1.string = "stu1"
	s1.int = 14
	fmt.Println(s1)

	var m1 myInt
	m1 = 100
	var m2 myInt = 99
	var m3 = myInt(98)
	m4 := myInt(97)
	fmt.Println(m1, m2, m3, m4)
	m4.hello()
}
