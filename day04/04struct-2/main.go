package main

import (
	"encoding/json"
	"fmt"
)

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

// 	结构体中可以嵌套其他结构体或结构体指针
type address struct {
	country string
	city    string
}

type detail struct {
	CEO       string
	startYear int64
}

type company struct {
	name    string
	address address // 正常嵌套
	detail          // 匿名嵌套
}

// 匿名嵌套结构体
// 语法糖，匿名嵌套时，可以直接访问嵌套的结构体的字段
// 正常嵌套访问字段 company1.address.city
// 匿名嵌套访问字段 company1.CEO ，可以省略匿名嵌套结构体
// 先在自己的结构体字段中找，然后再去匿名嵌套的结构体字段中找

// 匿名嵌套结构体字段冲突
// 如果两个匿名嵌套的结构体中有同名的字段，就会存在冲突，无法直接访问，要像访问正常嵌套结构体字段一样，指定匿名嵌套的结构体
// company1.detail.CEO

// 结构体的“继承”，模拟实现其他语言中面向对象的继承
// Go语言中没有继承的概念

type animals struct {
	name string
}

func (a *animals) move() {
	fmt.Printf("%s can move\n", a.name)
}

type dog struct {
	feet    int8
	animals // 通过嵌套匿名结构体实现继承，animals的字段和方法，dog都具有了
}

func (d *dog) bark() {
	fmt.Printf("%s can bark\n", d.name)
}

// 结构体中首字母大写的字段可公开访问，首字母小写的字段私有，仅在定义该结构体的包中可访问

// 结构体与JSON
// JSON，JavaScript Object Notion，轻量级的数据交换格式，易于阅读和编写，也易于机器解析和生成。
// 键值对组合中的键名写在前面并用双引号包裹，使用冒号分隔，然后紧跟值，多个键值之间用英文逗号分隔，键值之间和键值对之间没有空格

// 结构体变量转换成JSON格式的字符串，序列化
// encoding/json包
// 结构体中的字段名要首字母大写才可以，否则json包无法访问，就无法处理
// 结构体中首字母大写的字段可以被json包访问，处理输出
// 如果需要输出的键值对中key首字母要小写，就要用Tag结构体标签
// Tag是结构体的元信息，可以在运行的时候通过反射机制读取出来
// Tag在结构体字段的后方定义，由一对反引号包裹起来
// Tag必须严格遵守键值对的规则。结构体标签的解析代码的容错能力很差
// 一旦格式写错，编译和运行都不会提示任何错误，通过反射也无法正确取值，例如不要在key和value之间添加空格

// JSON格式的字符串转换成Go语言中能够识别的结构体变量，反序列化

type singer struct {
	Name   string `json:"name"` //通过指定tag实现json序列化该字段时的key
	Age    int8   // json序列化默认使用字段名作为key
	Album album
	gender string // 首字母小写的字段，私有不能被json包访问
}

type album struct {
	AlbumName   string
	ReleaseYear int64
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

	c1 := company{
		name: "QAX",
		address: address{
			country: "China",
			city:    "Beijing",
		},
		detail: detail{
			CEO:       "Qi",
			startYear: 2018,
		},
	}

	fmt.Println(c1)
	fmt.Println(c1.name, c1.address.city)
	fmt.Println(c1.CEO, c1.detail.startYear) // 匿名嵌套的结构体字段可以直接访问

	d1 := dog{
		feet: 4,
		animals: animals{
			name: "zero",
		},
	}

	d1.bark()
	d1.move()

	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Printf("marshal failed, error:%v\n", err)
		return
	}
	fmt.Printf("%s\n", data) //{"CEO":"Qi"}，只有CEO字段大写所以json能访问，才能获取，其他字段都没有办法处理
	fmt.Printf("%T\n",data) //[]uint8字节切片格式，不是字符串,格式化打印的时候用%s，或者string(data)转成字符串后再打印
	fmt.Println(string(data))

	singer1 := singer{
		Name: "Taeyeon",
		Age:  31,
		Album: album{
			AlbumName:   "I",
			ReleaseYear: 2015,
		},
		gender: "female",
	}

	data, err = json.Marshal(singer1)
	if err != nil {
		fmt.Printf("Marshal failed, error:%v\n", err)
		return
	}
	fmt.Printf("%v\n", string(data)) // {"name":"Taeyeon","Age":31,"Album":{"AlbumName":"I","ReleaseYear":2015}}
	// 通过Tag指定了json中Name字段key输出为name
	// 由于gender字段首字母小写，json包无法访问

	// 反序列化
	str := `{"name":"Kim","age":18}` // 反序列化不用在意key首字母大小写，但要与字段名相同，顺序也要一致，有tag的要跟tag对应
	var singer2 singer
	json.Unmarshal([]byte(str), &singer2) // 需要将字符串转换成字符切片才能传入，为了能修改singer2的变量需要传入指针
	fmt.Printf("%#v\n",singer2)
	str2 := `{"name":"Yonna","age":29,"album":{"albumname":"A Walk to Remember","releaseyear":2015}}`

	var singer3 singer
	json.Unmarshal([]byte(str2), &singer3)
	fmt.Printf("%#v\n", singer3)
}
