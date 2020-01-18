package main

import "fmt"

/*
内置函数
close 用来关闭channel
len 求长度（string、array、slice、map、channel）
new 分配内存，用于值类型，如int、stuct。返回的是指针
make 分配内存，用于引用类型，如chan、map、slice。返回的是引用类型自身
append 追加元素到数组、slice中

panic和recover
go目前没有异常机制，panic后程序崩溃退出，recover尝试恢复
panic可以在任何地方引发，但recover只有在defer调用的函数中有效
defer一定要在可能引发panic的语句之前定义

fmt包
Print系列函数会将内容输出到系统的标准输出
Print函数直接输出内容，采用默认格式将参数格式化并写入标准输出，如果两个相邻的参数都不是字符串，会在他们的输出之间添加空格
Printf函数支持格式化输出字符串，根据format参数生成格式化的字符串并写入标准输出
Println函数会在输出内容的结尾添加一个换行符，相邻参数的输出之间添加空格
func Print(a ...interface{}) (n int, err error)
func Printf(format string, a ...interface{}) (n int, err error)
func Println(a ...interface{}) (n int, err error)

*printf都支持format格式化参数

通用占位符，用于任何类型的变量
%v 值的默认格式输出，value
%+v 类似%v，但输出结构体时会添加字段名
%#v 值的Go语法表示，如输出字符串会带上双引号这种语法特征
%T 打印值的类型
%% 百分号，格式化字符串中%是特殊字符，也不能用转义符，用%%转义输出一个%

布尔型
%t true或false

整型
%d 十进制
%b 二进制
%o 八进制
%x 十六进制，使用a-f
%X 十六进制，使用A-F
%c 该值对应的unicode码值
%q 该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示


浮点数和复数
%f和%F 又小数部分，但无指数部分
%e和%E 科学计数法
%g和%G 根据实际情况选择上两种格式，以获得更简洁、更准确的输出

字符串和[]byte
%s 直接输出字符串或[]byte
%x 每个字节用两字符十六进制数表示（a-f）
%X 每个字节用两字符十六进制数表示（A-F）
%q 该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示

指针
%p 十六进制，并加上前导0x

宽度通过紧跟在百分号后面的十进制数指定，未指定宽度则表示值时除必须之外不做填充
精度（可选）通过宽度后跟点号再跟的十进制数指定
%f 默认宽度，默认精度
%9f 宽度9 默认精度
%.2f 默认宽度，精度2 // 点之前没有数字，表示默认宽度
%9. 宽度9，精度0 // 点之后没有数字，表示精度为0
%9.2 宽度9，精度2

*/
func funcA() {
	fmt.Println("func A")
}

func funcB() {
	defer func() {
		err := recover()
		fmt.Println(err)
		// 如果程序出现了panic错误，可以通过recover恢复
		if err != nil {
			fmt.Println("recover in func B")
		}
	}()
	panic("panic in func B")
}

func funcC() {
	fmt.Println("func C")
}

func main() {
	funcA()
	funcB()
	funcC()

	// 获取用户输入
	// Scan函数简单的以空格作为输入数据的分隔符
	//var s string
	//fmt.Scan(&s) // 值类型传参修改的是副本，要修改原本需要传入指针
	//fmt.Println("用户输入的内容是：", s)

	var (
		name  string
		age   int
	)
	//Scanf为输入数据指定了具体的输入内容格式，只有按照格式输入数据才会被扫描并存入相应变量
	fmt.Scanf("姓名：%s 年龄：%d\n", &name, &age)
	fmt.Println(name, age)

	// 遇到换行才停止扫描，最后一个数据之后必须有换行，遇到回车就结束扫描
	//fmt.Scanln(&name, &age)
	//fmt.Println(name, age)
}
