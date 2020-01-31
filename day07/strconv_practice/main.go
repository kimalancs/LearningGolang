package main

import (
	"fmt"
	"strconv"
)

// strconv包实现了基本数据类型和其字符串表示之间的相互转换

func main() {
	i := 2316
	ret := string(i) // 此处并不是int类型转换为string类型，而是把这个int变量当作ASCII码，来取对应的字符
	fmt.Println(ret) // ऌ

	iStr := strconv.FormatInt(int64(i), 10) // 数字转换成字符串，传入int64类型
	fmt.Printf("%T %#v\n", iStr, iStr)

	i2, _ := strconv.ParseInt(iStr, 10, 64) // 从字符串中解析出数字, 字符串中必须只有数字，返回的是int64，接受正负号
	i3, _ := strconv.Atoi(iStr)             // 字符串中解析数字，返回int类型
	iStr2 := strconv.Itoa(i)                // 数字转换成字符串，传入int类型
	fmt.Printf("%T %#v\n", i2, i2)
	fmt.Printf("%T %#v\n", i3, i3)
	fmt.Printf("%T %#v\n", iStr2, iStr2)
	bStr := strconv.FormatBool(true)
	fmt.Printf("%T %#v\n", bStr, bStr)
	b, _ := strconv.ParseBool(bStr)
	fmt.Printf("%T %#v\n", b, b)
	f := 1.23
	fStr := strconv.FormatFloat(f,'f', -1, 64) // 第三个参数prec控制精度，小数点后的位数，如果prec 为-1，则代表使用最少数量的、但又必需的数字来表示f
	fmt.Printf("%T %#v\n",fStr,fStr)
	f2, _ := strconv.ParseFloat(fStr, 64) // 字符串中解析浮点数
	fmt.Printf("%T %#v\n",f2, f2)
}
