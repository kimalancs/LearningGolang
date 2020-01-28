package main

import (
	"encoding/json"
	"fmt"
)

// 反射主要了解原理，写代码时很少自己写
// 但是很多包、框架，其原理来自反射，如json格式化

// 变量分为两部分
// 类型信息，预先定义好的元信息
// 值信息，程序运行过程中可动态变化的

// 程序在编译时，变量名被转换为内存地址，变量名不会被编译器写入到可执行部分
// 在运行程序时，程序无法获取自身的信息

// 反射是指在程序运行期，对程序本身进行访问和修改的能力

// 支持反射的语言可以在程序编译期将变量的反射信息，如字段名称、类型信息、结构体信息等整合到可执行文件中
// 并给程序提供接口访问反射信息，这样就可以在程序运行期获取类型的反射信息，并且有能力修改它们

// Go程序在运行期使用reflect包访问程序的反射信息

// 空接口可以存储任何类型的变量，如何获知空接口保存的数据是什么
// 反射就是在运行时动态的获取一个变量的类型信息和值信息

// 任何接口值都是由一个具体类型和具体类型的值两部分组成的
// 任何接口值在反射中都可以理解为由reflect.Type和reflect.Value两部分组成
// 并且reflect包提供了reflect.TypeOf和reflect.ValueOf两个函数来获取任意对象的Type和Value



type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	str := `{"name":"John","age":22}`
	var p person
	json.Unmarshal([]byte(str), &p) // 真正执行时
	fmt.Println(p.Name, p.Age)

}
