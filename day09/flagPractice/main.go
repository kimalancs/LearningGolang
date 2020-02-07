package main

import (
	"flag"
	"fmt"
	"time"
)

//flag包支持的命令行参数类型
// bool、int、int64、uint、uint64、float float64、string、duration


// flag.Args()  ////返回命令行参数后的其他参数，以[]string类型
// flag.NArg()  //返回命令行参数后的其他参数个数
// flag.NFlag() //返回使用的命令行参数个数

func main() {
	// 定义命令行参数方式1 
	// flag.Type(flag名, 默认值, 帮助信息)*Type 
	// 例如我们要定义姓名、年龄、婚否三个命令行参数，我们可以按如下方式定义：
	// name := flag.String("name", "张三", "姓名")
	// age := flag.Int("age", 18, "年龄")
	// married := flag.Bool("married", false, "婚否")
	// delay := flag.Duration("d", 0, "时间间隔")
	// 需要注意的是，此时name、age、married、delay均为对应类型的指针


	//定义命令行参数方式2
	// flag.TypeVar(Type指针, flag名, 默认值, 帮助信息)
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")

	//解析命令行参数
	// 调用flag.Parse()来对命令行参数进行解析
	// -flag xxx （使用空格，一个-符号）
	// --flag xxx （使用空格，两个-符号）
	// -flag=xxx （使用等号，一个-符号）
	// --flag=xxx （使用等号，两个-符号）
	// 其中，布尔类型的参数必须使用等号的方式指定。
	// Flag解析在第一个非flag参数（单个”-“不是flag参数）之前停止，或者在终止符”–“之后停止。
	flag.Parse()
	fmt.Println(name, age, married, delay)
	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
}
