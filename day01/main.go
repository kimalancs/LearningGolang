// main包才能生成可执行文件，其他包会打包到main包的可执行文件中。
package main

// 导入语句，双引号
import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

// 函数外只能放置标识符（变量、常量、函数、类型）,函数外每个语句必须以关键字开始（var、const、func等）
// 标识符由字母数字和下划线组成，不能以数字开头，25个关键字，37个保留字
// 变量，数据存储在内存空间，通过变量名找到内存中对应的变量，变量替代内存地址可读性更强，计算机操作还是通过内存地址，源代码编译成可执行文件是变量名会换成内存地址
// 变量需声明再使用，var s1 string 静态类型，必须指定类型，类型决定使用内存的大小
// 支持批量声明
// var name string
// var age int
// var isOK bool
// 全局变量声明后可以不使用，因为可能在其他包中被使用
var (
	name string // 不指定变量值，初始化为对应类型的空值，“”
	age  int    // 0
	isOK bool   // false
)

const pi = 3.1415 // 程序运行期间恒定不变的量，常量通常定义在全局
const e = 2.7182

const (
	statusOK = 200 //批量声明
	notFound = 404
)

const (
	n1 = 100 // 批量声明常量时，如果某一行为赋值，默认该常量与上一行的常量值相同
	n2
	n3 = 200
	n4 = iota // 3 iota在cosnt关键字出现时将被重置为0，const中每新增一行，iota加一
	n5        //4 iota可用于实现类似枚举
	_
	n6 //6
	n7 = 100
	n8 = iota // 8
)

const (
	d1, d2 = iota + 1, iota + 2 // 1，2 多个常量声明在一行，仍然是每增一行，iota加一
	d3, d4 = iota + 1, iota + 2 // 2，3
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func foo() (int, string) {
	return 10, "Qian"
}

// 整形分两类，带符号和无符号
// 有符号整形
// int8 	-128到127
// int16 	-32768到32767
// int32 	-2147483648到2147483647  就是rune
// int64 	-9223372036854775808到9223373036854775807
// 无符号整形
// uint8 	0到255
// uint16 	0到65535
// uint32 	0到4294967295
// uint64 	0到18446744073709551615
// uint8 就是 byte
// int16 对应C语言中的short型
// int64 对应C语言中的long型

// 写跨平台应用时要注意
// uint 32位的操作系统上就是uint32， 64位操作系统上就是uint64
// int 32位的操作系统上就是int32， 64位操作系统上就是int64

// uintptr 无符号整形，用于存放一个指针

// 浮点数，小数
// float32
// float64

// 布尔类型，bool
// true
// false 默认值为false
// 不允许将整型强制转换为布尔类型，python可以
// 无法参加数值运算，也无法与其他类型进行转换

//复数
//complex128
//comlpex64

// 字符串只能用双引号包裹

// 单引号包裹的是字符，一个单独的字符（字母、数字、符号、汉字等）

// Go语言中字符分两种
// uint8类型，或者叫byte型，代表ASCII码的一个字符，比如英文字母
// rune类型，代表一个UTF-8字符，比如中文、日文等
// 当需要处理中文、日文或者其他符合字符时，需要使用rune类型，rune类型实际是一个int32
// 字符串底层是一个byte数组，可以和[]byte类型相互转换
// 字符串是不能修改的。字符串是由byte字节组成，所以字符串的长度是byte字节的长度
// rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成

func traversalString() {
	s := "hello 北京"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i]) // %c 字符
	}
	// UTF-8编码下一个中文汉字由3～4个字节组成，不能直接按照字节去遍历一个包含中文的字符串，常见中文一般3个字节
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

//修改字符串
//需要将其转换成[]rune或[]byte，完成后再转换成string，无论哪种转换，都会重新分配内存，并复制字节数组

func changeString() {
	s1 := "big"
	//强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '胡'
	fmt.Println(string(runeS2))
}

// 只有强制类型转换，没有隐式类型转换，且只能在两个类型之间支持相互转换时使用
// T(表达式)
// T代表要转换的类型，表达式包括变量，复杂算子和函数返回值等

func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b))) // 赋值给c时又要转换回int
	fmt.Println(c)
}

// 统计 "hello 北京"字符串中汉字的个数
func getHanNumber(str string) float64 {
	var total float64

	for _, r := range str {
		if unicode.Is(unicode.Han, r) {
			total = total + 1
		}
	}
	return total
}


//程序的入口函数，没有参数，没有返回值
func main() {
	name = "allen"
	age = 16
	isOK = true

	//Go语言中非全局变量变量声明了必须要使用，否则无法编译成功，无效语句会导致可执行文件变大
	// 驼峰式命名
	// var student_name string //不推荐
	var studentName string //推荐用小驼峰
	//var	StudentNmae string //大驼峰
	fmt.Print(isOK)             // 输出要打印的内容，不换行
	fmt.Println("Hello world!") // fmt.Println 打印完指定内容后会自动加一个换行符
	fmt.Printf("Name:%s", name) // fmt.Printf 格式化输出，%s 占位符，s string 字符串，使用name变量的值去替换占位符，不换行

	studentName = "kim"    //同一个作用域中不能重复声明同名的变量
	fmt.Print(studentName) // 非全局变量一定要使用

	var s1 = "SNSD" // 类型推导，根据值自动判断类型
	s2 := "Taeyeon" // 短变量声明，直接赋值，只能用在函数内部
	fmt.Println(s1 + s2)

	x, _ := foo() //匿名变量，anonymous variable，用一个下划线，在函数返回多个值，而你不需要其中某个时，用于占位，忽略某个值
	_, y := foo() // 匿名变量不占用命名空间，不分配内存，不存在重复声明，lua中也被称为哑元变量
	fmt.Println("x=", x)
	fmt.Println("y=", y)

	fmt.Println("n1", n1)
	fmt.Println("n2", n2)
	fmt.Println("n4", n4)
	fmt.Println("n5", n5)
	fmt.Println("n6", n6)
	fmt.Println("n8", n8)
	fmt.Println("d1", d1)
	fmt.Println("d2", d2)
	fmt.Println("d3", d3)
	fmt.Println("d4", d4)

	var a int = 77        // 无法直接定义二进制数，先定义其他十进制数，转换输出二进制数
	fmt.Printf("%d\n", a) // d 十进制
	fmt.Printf("%b\n", a) // b 二进制
	fmt.Printf("%o\n", a) // o 八进制
	fmt.Printf("%x\n", a) // x 十六进制

	// 可以定义八进制和十六进制
	b := 0115 // 0开头代表八进制，给文件分配权限时使用八进制
	c := 0x4d // 0x开头代表十六进制，涉及内存地址时使用十六进制
	fmt.Printf("%d\n", b)
	fmt.Printf("%d\n", c)

	// %T 查看变量的类型
	fmt.Printf("%T\n", c)

	// 声明int8类型的变量
	d := int8(9) // 明确指定int8类型，否则就是int类型
	fmt.Printf("%T\n", d)

	fmt.Println(math.MaxFloat32)
	// math.MaxFloat64
	f1 := 1.23456
	fmt.Printf("%T\n", f1) // 默认都是float64类型
	f2 := float32(1.23456) // 必须明确指定float32才能声明，显式声明
	fmt.Printf("%T\n", f2)
	//f1不能赋值给f2，类型不同

	b1 := true
	var b2 bool
	fmt.Printf("%T value:%v\n", b1, b1) // %v 查看变量的值，什么类型都可以用，%#v 会添加类型的特征，比如字符串会带“”
	fmt.Printf("%T value:%v\n", b2, b2)

	// \ 转义符，将字符串中本来有特殊含义的，比如双引号包裹字符串，字符串中有双引号的，使用转义符来表示该双引号不具特殊意义
	// \r 回车符
	// \n 换行符
	// \t 制表符
	// \' 单引号
	// \" 双引号
	// \\ 反斜杠，比如Windows路径时使用

	path := "\"D:\\Go\\src\\studygo\\\""
	fmt.Println(path)

	// 多行字符串，使用反引号` 英文字符模式下ESC键下方
	s3 := `
		世情薄
		人情恶
		雨送黄昏花易落
	`
	fmt.Println(s3)

	s4 := `D:\Go\src\stuygo\` // 反引号包裹起来的字符串原样输出，只有换行符还有效，其余转义字符无效，类似python中的raw字符串
	fmt.Println(s4)

	// 字符串拼接
	// len(str) 求长度
	// +或fmt.Sprintf 拼接字符串
	// strings.Split 分割字符串
	// strings.contains 判断是否包含
	// stirngs.HasPrefix，strings.HsaSuffix 前缀\后缀判断
	// strings.Index()，Strings.LastIndex() 子串出现的位置
	// strings.Join(a[]str,sep string) join操作

	fmt.Println(len(s4))
	name := "Kim"
	word := " is a good guy"
	s5 := name + word
	fmt.Println(s5)
	s6 := fmt.Sprintf("%s%s", name, word)
	fmt.Println(s6)

	s7 := strings.Split(s4, "\\") // 把字符串按照第二个参数sep分割成多个子串，输出一个切片
	fmt.Println(s7)
	fmt.Printf("%T\n", s7)

	fmt.Println(strings.Contains(s5, "is OK"))
	fmt.Println(strings.Contains(s5, "good guy"))

	fmt.Println(strings.HasPrefix(s5, "Kim"))
	fmt.Println(strings.HasSuffix(s5, "good"))

	fmt.Println(strings.Index(s5, "good"))

	fmt.Println(strings.Join(s7, "\\"))

	fmt.Println(len(s7))

	traversalString()

	changeString()

	sqrtDemo()

	n := 10
	var f3 float64
	f3 = float64(n)
	fmt.Println(f3)
	fmt.Printf("%T\n", f3)

	// 流程控制
	//if条件判断
	age := 18
	if age >= 18 {
		fmt.Println("欢迎光临xx网咖！")
	} else {
		fmt.Println("未成年禁止入内！")
	}

	// 多个条件
	if age := 38; age > 35 { //此时age是局部变量，只在这个if语句的作用域中生效，与上一段的age不冲突，执行完此if语句，变量就销毁，减少内存占用
		fmt.Println("人到中年")
	} else if age >= 18 {
		fmt.Println("大好年华")
	} else {
		fmt.Println("好好学习")
	}

	fmt.Println(age) // 18

	/*
		for循环，只有这一种循环。多种变种。
		基本格式
		for 初始语句；条件表达式；结束语句{
			循环体语句
		}
		条件表达式返回true时，循环体不停进行循环，再执行结束语句，然后再进行条件判断，直到条件表达式返回false时自动推出循环
	*/
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//初始语句和结束语句都可以省略
	//变种1，省略初始语句
	var i = 5
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	/*
		变种2，省略结束语句
		for i < 10 { // 也可以写成 for ; i < 10; {
			fmt.Println(i)
			i++
		}

		变种3，无限循环，死循环
		for {
			fmt.Println("123")
		}
	*/

	// for range 循环
	// 用于遍历数组、切片、字符串、map、通道
	// 返回索引和值
	s8 := "Hello 山东"
	for i, v := range s8 {
		fmt.Println(i, v)
		fmt.Printf("%d %c\n", i, v)
	}

	fmt.Println(getHanNumber("hello 北京"))

	// 打印99乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, j*i)
		}
		fmt.Println()
	}
	fmt.Println() // 换行

	for i := 1; i < 10; i++ {
		for j := i; j < 10; j++ {
			fmt.Printf("%d*%d=%d\t", i, j, j*i)
		}
		fmt.Println()
	}
	fmt.Println()

	for i := 9; i >= 1; i-- {
		for j := i; j >= 1; j-- {
			fmt.Printf("%d*%d=%d\t", j, i, j*i)
		}
		fmt.Println()
	}
	fmt.Println()

	// break 跳出for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("over")

	// continue 跳过此次循环，继续下一轮循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("over")

	//switch case 条件判断
	// 简化大量的if else判断（一个变量和具体的值做比较）

	finger := 5

	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default: // 每个switch只能有一个default分支
		fmt.Println("无效的输入")
	}

	switch n := 7; n {
	case 1, 3, 5, 7, 9: // 一个分支可以有多个值，用英文逗号分隔
		fmt.Println("奇数")
	case 2, 4, 6, 8, 10:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}

	switch { // 分支可以使用表达式，此时switch语句后不需要跟判断变量
	case age < 18:
		fmt.Println("好好学习")
	case age >= 18 && age < 25:
		fmt.Println("享受青春")
	case age >= 25 && age <= 65:
		fmt.Println("努力工作")
	default:
		fmt.Println("珍惜时光")
	}

	// fallthrough可执行满足条件的case的下一个case，为兼容C语言中的case设计，不要使用
	string1 := "a"
	switch {
	case string1 == "a":
		fmt.Println("a")
		fallthrough
	case string1 == "b":
		fmt.Println("b")
	case string1 == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}

	// goto 跳转到指定标签(不鼓励使用，比较晦涩)
	// 通过标签进行代码间的无条件跳转
	// 在快速跳出循环、避免重复退出上有一定的帮助
	// 简化一些代码的实现过程，如双层嵌套的for循环要退出时
	/* var breakFlag bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				breakFlag = true
				break // 跳出内层循环
			}
			fmt.Println(i, j)
		}
		//外层for循环判断
		if breakFlag {
			break // 跳出外层循环
		}
	}
	*/

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Println(i, j)
		}
	}
	return
	// 标签 Label
	breakTag:
		fmt.Println("end")

}
