package main

import "fmt"

type student struct {
	name string
	age  int8
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "kim", age: 20},
		{name: "tae", age: 19},
		{name: "yonna", age: 18},
	}

	for _, stu := range stus {
		m[stu.name] = &stu // 三次传入的内存地址都一样
		fmt.Printf("%p\n",&stu)
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
	fmt.Println(m)
}
