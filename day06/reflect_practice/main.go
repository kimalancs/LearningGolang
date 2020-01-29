package main

import (
	"encoding/json"
	"fmt"
	"reflect"
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

// 在反射中类型还划分为两种，类型Type、种类Kind
// Kind是指底层的类型，比如结构体这种大类，通过.Name()获取
// Type是指具体的类型，比如自定义的struct类的person类型，通过.Kind()获取
// 数组、切片、Map、指针等类型的变量，正常的它们的.Name()和.Kind()都一样，基于这几个类型来自定义的类型，它们的.Name()是自定义的类型名，而.Kind()是其基于的类型

// 容易混淆的地方
// TypeOf()和ValueOf()都有.Kind()方法，一个是类型的底层类型，一个是值的类型

// reflect.ValueOf()返回的是reflect.Value类型，其中包含了原始值的值信息
// reflect.Value与原始值之间可以互相转换
// 获取原始值的方法
// Int() int64
// Uint() uint64
// Float() float64
// Bool() bool
// Bytes() []byte
// String() string
// Interface() interface{}

// 通过反射修改变量的值
// 在函数中修改变量的值必须传入指针
// 反射中使用专有的Elem()方法获取指针的值

// func(v Value) IsNil() bool
// IsNil()会报告v持有的值是否为nil，v持有的值的分类必须是通道、函数、接口、映射、指针、切片之一，否则IsNil方法会导致panic
// 常被用于判断指针是否为空

// func(v Value) IsVaild() bool
// IsVaild()返回v是否持有一个值。如果v是Value零值则会返回false，此时v除了IsValid、String、Kind之外的方法都会导致panic
// 常被用于判定返回值是否有效


// 结构体反射
// 任意值通过reflect.TypeOf()获得反射对象信息后，如果它的类型是结构体，可以通过反射值对象（reflect.Type）的NumField()和Field()方法获取结构体成员的详细信息
// Field(i init) StructField 根据索引，返回索引对应的结构体字段
// NumField() int 返回结构体成员字段数量
// 这两个结合，可以先得到字段数量，然后遍历获取所有字段

// FieldByName(name string) (StructField, bool) 根据给定字符串返回字符串对应的结构体字段的信息
// FieldByIndex(index []int) StructField 多层成员访问时，根据[]int提供的每个结构体的字段索引，返回字段的信息
// FieldByNameFunc(match func(string) bool) (StructField, bool) 根据传入的匹配函数匹配需要的字段
// NumMethod() int 返回该类型的方法集中方法的数量
// Method(int) Method 返回该类型方法集中的第i个方法
// MethodByName(string) (Method, bool) 根据方法名返回该类型方法集中的方法

// StructField类型用来描述结构体中的一个字段的信息
// type StructField struct {
//	// Name是字段的名字，PkgPath是非导出字段的包路径，对导出字段该字段为""
//	Name string
//	PkgPath string
//	Type 	Type		// 字段的类型
//	Tag 	StructTag	// 字段的标签
//	Offset	uintptr		// 字段在结构体中的字节偏移量
//	Index	[]int		// 用于Type.FieldByIndex时的索引切片
//	Anonymous bool		// 是否匿名字段
//}


// 反射是一个强大并富有表现力的工具，能让我们写成更灵活的代码
// 但是不能滥用，原因如下：
// 基于反射的代码极其脆弱，反射中的类型错误会在真正运行时才引发panic，可能是代码写完很久之后才发现
// 大量使用反射的代码通常难以理解
// 反射的性能低下，基于反射实现的代码通常比正常代码运行速度慢一到两个数量级



type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", t)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
}

func reflectValue(x interface{}){
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取的整形的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Int()从反射中获取的整形的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Int()从反射中获取的整形的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200)
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func (p person) Study() string {
	msg := "好好学习，天天向上"
	return msg
}

func (p person) Sleep() string {
	msg := "好好睡觉，健康成长"
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	fmt.Println(v.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是[]reflect.Value类型
		var args []reflect.Value
		v.Method(i).Call(args)
	}
}

func main() {
	str := `{"name":"John","age":22}`
	var p person
	json.Unmarshal([]byte(str), &p) // 接收的第二个参数是空接口，真正执行时才会确定变量类型
	fmt.Println(p.Name, p.Age)

	var a float32 = 1.2
	reflectType(a)
	var b string = "hello"
	reflectType(b)
	reflectType(p)

	type myint int64

	var c myint = 11
	reflectType(c)

	var d float32 = 3.14
	var e int64 = 100
	reflectValue(d)
	reflectValue(e)
	// 将int类型的原始值转换成reflect.Value类型
	f := reflect.ValueOf(10)
	fmt.Printf("type:%T\n", f)

	//reflectSetValue(e) // 传入的是副本，reflect包会引发panic
	reflectSetValue2(&e)
	fmt.Println(e)


	var g *int
	fmt.Println("g IsNill:", reflect.ValueOf(g).IsNil())

	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())

	h := struct{}{}
	// 尝试从结构体中查找"abc"字段
	fmt.Println("不存在的结构体成员：", reflect.ValueOf(h).FieldByName("abc").IsValid())
	// 尝试从结构体中查找"abc"方法
	fmt.Println("不存在的结构体方法：", reflect.ValueOf(h).MethodByName("abc").IsValid())

	i := map[string]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("map中不存在的键：", reflect.ValueOf(i).MapIndex(reflect.ValueOf("kim")).IsValid())


	stu := person {
		Name: "taeyoen",
		Age: 18,
	}

	t := reflect.TypeOf(stu)
	fmt.Println(t.Name(), t.Kind())
	// for循环遍历结构体中所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}
	// 通过字段名获取指定结构体字段信息
	if ageField, ok := t.FieldByName("Age"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", ageField.Name, ageField.Index, ageField.Type, ageField.Tag.Get("json"))
	}

	printMethod(stu)
}
