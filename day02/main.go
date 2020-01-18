package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// 单行注释

/*
	多行注释
*/

/*
运算符

算数运算符
+ 加
- 减
* 乘
/ 除
% 取模，求余

单独的语句，不能放在=的右边赋值
++ 自增
-- 自减

关系运算符，返回一个布尔值，true或false
== 相等，判断是否相同，相同的类型才能比较
!= 不等于
> 大于
< 小于
>= 大于等于
<= 小于等于

逻辑运算符
&& 逻辑AND运算符，与，两边都真才为真，又一个假就是假
|| 逻辑OR运算符，或，两边又一个真就是真，两个都假才是假
! 逻辑NOT运算符，非，取反，原来为真就为假，原来为假就为真

位运算符，二进制位
& 按位与，参与运算的两个数对应的二进制位相与，两位为1才为1
5的二进制位101
2的二进制位010
5 & 2 // 000

| 按位或，参与运算的两个数对应的二进制位相或，两位有一个为1就为1
5 ｜ 2 // 111

^ 按位异或，参与运算的两个数对应的二进制位相异或，当两对应的二进制位相异时（两个不一样），结果为1
5 ^ 2 // 111

<< 左移n位就是乘以2的n次方
a << b 把a的各二进制位全部左移b位，高位丢弃，低位补0
5 << 2 // 10100 5*4=20

>> 右移n位就是除以2的n次方
a >> b 把a的各二进制位全部右移b位
5 >> 2 // 001 1

赋值运算符
= 赋值
+= 相加后再赋值
-+ 相减后再赋值
*= 相乘后再赋值
/= 相除后再赋值
%= 取余后再赋值
<<= 左移后后再赋值
>>= 右移后再赋值
&= 按位与后再赋值
|= 按位或后再赋值
^= 按位异或后再赋值

x += 1 // x = x + 1
x -= 1 // x = x - 1
x *= 1 // x = x * 1
x /= 1 // x = x / 1
x %= 1 // x = x % 1

*/

func main() {
	// array 数组，容器，同一种数据类型元素的集合，声明时就要确定类型和长度，可以修改数组成员，但是长度不可变
	var a1 [3]bool
	var a2 [4]bool
	fmt.Printf("a1:%T a2:%T\n", a1, a2)

	// 数组的初始化
	// 不指定，默认元素为零值，如布尔值数组，默认全是false
	a2 = [4]bool{true, false, true, false}
	fmt.Println(a1, a2)

	a3 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	a4 := [...]int{1, 2, 3}  //根据初始化值推算数组长度
	a5 := [5]int{1, 2}       //可以只初始化一部分元素
	a6 := [5]int{0: 1, 4: 2} // 可通过索引指定初始化某个元素
	fmt.Println(a3, a4, a5, a6)

	// 数组的遍历
	// 根据索引遍历
	citys := [...]string{"北京", "上海", "深圳"}
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	// for range遍历
	for i, v := range citys {
		fmt.Println(i, v)
	}

	// 多维数组
	var a7 [3][2]int
	a7 = [3][2]int{
		{0, 1},
		{1, 3},
		{2, 4},
	}
	fmt.Println(a7)

	a8 := [...][2]string{
		{"beijing", "shanghai"},
		{"NewYork", "London"},
	}
	fmt.Printf("%T %v\n", a8, a8)
	fmt.Println(a8[1][1])

	// 多维数组的遍历
	for _, v := range a8 {
		for _, v := range v {
			fmt.Printf("%s\t", v)
		}
		fmt.Println()
	}

	// 数组是值类型，赋值和传参会复制整个数组，因此改变副本的值，不会改变本身的值

	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 10
	fmt.Println(b1, b2)

	// 数组支持 == 和 != 操作符，因为内存总是被初始化过的
	// [n]*T表示指针数组
	// *[n]T表示数组指针

	// 练习题1，求数组[1, 3, 5, 7, 8]所有元素的和
	c1 := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, v := range c1 {
		sum = sum + v
	}
	fmt.Println(sum)

	// 练习题2，找出数组中和为指定值的两个元素的下标，比如[1, 3, 5, 7, 8]中找到和为8的两个元素的下标分别为(0, 3)和(1, 2)
	for i := 0; i < len(c1); i++ {
		for j := i + 1; j < len(c1); j++ {
			if c1[i]+c1[j] == 8 {
				fmt.Printf("(%d, %d)\n", i, j)
			}
		}
	}

	// slice 切片
	// 数组的长度锁定且数组长度属于类型的一部分，所以数组有很多局限性
	// 拥有相同类型元素的可变长度序列
	// 基于数组类型做的一层封装，底层就是一个数组
	// 非常灵活，支持自动扩容
	// 引用类型，内部结构包括地址、长度、容量
	// 一般用于快速地操作一块数据集合
	// var name []T
	// name表示变量名
	// T表示切片中元素类型
	// 与数组的区别，不限制长度
	var s1 []string              // 未初始化，仍然为空 nil，相当于没有开辟内存空间,nil值的切片没有底层数组
	var s2 = []bool{true, false} // 已初始化，建立了底层数组
	s3 := []int{}                // 已初始化，没有元素，长度为0，但不建立了底层数组，不是nil
	fmt.Println(s1, s2, s3)
	fmt.Printf("%T, %T, %T\n", s1, s2, s3)
	// 切片是引用类型，不支持直接比较，只能和nill比较(不能用s1 == s2)
	// 一个nil值的切片的长度和容量都为0，但长度和容量都为0的切片不一定是nil
	// 要判断一个切片是否为空，用len(s) == 0来判断，不应该使用 s == nil
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	fmt.Println(s3 == nil)

	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))
	fmt.Printf("len(s3):%d cap(s3):%d\n", len(s3), cap(s3))

	// 由数组得到切片
	d1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	d2 := d1[0:4] // 基于一个数组切割，左包含右不包含，左闭右开，包含索引0的元素，不包含索引4的元素
	d3 := d1[2:]  // [2:len(d1)]
	d4 := d1[:6]  // [0:6]
	d5 := d1[:]   // [0:len(d1)]
	fmt.Println(d2, d3, d4, d5)
	// 切片指向一个底层的数组
	// 切片的长度就是元素的个数
	//切片的容量是指底层数组从切片的第一个元素到底层数组最后的元素数量
	// 真正的数据保存在底层数组中，切片只有指针、长度、容量
	fmt.Printf("len(d2):%d cap(d2):%d\n", len(d2), cap(d2)) //d2从底层数组最开始开始切
	fmt.Printf("len(d3):%d cap(d3):%d\n", len(d3), cap(d3)) //d3从底层数组索引2开始切

	// 切片再切片，索引不能超出原数组的长度，否则会出现索引越界的错误
	d1[12] = 15 // 切片是引用类型，从d1产生的切片都指向了这个底层数组，数组元素变了，切片的对应元素也变了
	d6 := d3[3:]
	fmt.Println(d3, d6)
	fmt.Printf("len(d6):%d cap(d6):%d\n", len(d6), cap(d6))

	// make()函数构造切片，make([]T, size, cap) 动态创建一个切片,T元素类型，size元素数量（长度），cap 容量
	e1 := make([]int, 2, 10)
	fmt.Println(e1)
	fmt.Printf("len(e1):%d cap(e1):%d\n", len(e1), cap(e1))

	// 切片的赋值拷贝
	// 拷贝前后两个变量共享底层数组，对一个切片的修改会影响另一个切片的内容

	e2 := e1
	e2[0] = 1
	fmt.Println(e1, e2)

	// 切片的遍历
	// 索引遍历
	for i := 0; i < len(e1); i++ {
		fmt.Println(i, e1[i])
	}
	// for range遍历
	for index, value := range e1 {
		fmt.Println(index, value)
	}

	// append()函数可为切片动态追加元素，每个切片会指向一个底层数组，这个数组的容量够用就添加新增元素
	// 当底层数组不能容纳新增的元素时，切片会自动按照一定的策略扩容，这是该切片指向的底层数组会替换
	// 扩容操作通常发生在append()函数调用时，通常需要用原变量接收append函数的返回值
	// 必须用原来的切片变量接收返回值
	// 可以一次追加多个元素
	// 容量的扩容策略，早期就是翻倍，后来进行了优化
	// 先判断新申请容量大于2倍的旧容量，最终容量就是新申请容量
	// 再判断，如果旧切片长度小于1024，则最终容量就是旧容量的2倍
	// 否则判断，旧切片长度大于1024，则最终容量从旧容量开始循环增加原来的1/4，直到最终容量大于等于新申请的容量
	// 如果最终容量计算值溢出，则最终容量就是新申请容量
	// 扩容策略根据切片中元素类型不同而做不同的处理，int和string类型处理方式就不一样
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v len:%d cap:%d ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}

	var citySlice []string
	citySlice = append(citySlice, "广州", "北京", "上海", "深圳")
	fmt.Println(citySlice)
	fmt.Printf("%v len:%d cap:%d\n", citySlice, len(citySlice), cap(citySlice))

	citysAppend := []string{"成都", "杭州", "嘉兴"}
	citySlice = append(citySlice, citysAppend...) // 不能直接把切片追加到切片，用...把切片拆开成多个元素就可以
	fmt.Printf("%v len:%d cap:%d\n", citySlice, len(citySlice), cap(citySlice))

	// copy函数复制切片，可快速将一个切片的数据复制到另一个切片空间中
	// copy(destSlice, srcSlice []T)
	// destSlice目标切片 srcSlice数据来源切片
	// 将一个切片赋值给另一个切片，由于引用类型，共享一个底层数组，一个修改，另一个也受影响
	// copy函数复制切片，不再是同一个底层数组，一个修改，另一个不受影响

	f1 := []int{1, 2, 3}
	f2 := f1 // 赋值
	var f3 = make([]int, 3, 3)
	copy(f3, f1) //copy复制切片
	f1[0] = 5    // 修改f1后，f2同样修改，f3不变
	fmt.Println(f1, f2, f3)

	// 没有删除切片元素专用方法，可以使用切片本身的特性来删除元素
	// 删除索引为index的元素，操作方法 a = append(a[:index],a[index+1:]...)
	// 把切片再次切片出包含index索引元素前的所有元素和包含index元素后所有元素的两个切片，再把后者追加到前者中
	g := [...]int{1, 2, 3}
	g1 := g[:]
	g1 = append(g1[:1], g1[2:]...) // append修改了底层数组
	fmt.Println(g, g1)

	// 练习题
	var h = make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		h = append(h, i)
	}
	fmt.Println(h)

	// 指针 pointer
	// 不存在指针操作，更安全，C语言允许操作指针，效率高，但是对程序员要求高
	// 数据载入内存后，在内存都有对应的地址，就是指针
	// 保存一个数据在内存中的地址，就需要指针变量
	// GO语言指针不能进行偏移和运算，只需要记住两个符号
	// & 取地址操作符
	// * 取值操作符，根据地址取出地址指向的值
	// 一对互补操作符
	// 对变量进行取地址&操作，可以获得这个变量的指针变量
	// 指针变量的值是指针地址
	// 对指针变量进行取值*操作，可以获得指针变量指向的原变量的值
	// 内存地址就是个十六进制的数
	// 每种值类型都有对应的指针类型，如*int，*string等
	i1 := 18
	fmt.Println(&i1)
	i2 := &i1
	fmt.Printf("%T %v\n", i2, i2)
	i3 := *i2
	fmt.Printf("%T %v\n", i3, i3)

	// 指针传值
	// 将变量传入内存，是复制，函数内的修改不影响原变量
	// 将指针传入内存，传入的是指向某变量的内存地址，函数内的修改会影响到该变量

	// new和make
	// 下面的代码可以编译，但是运行时会出现panic，引用类型的变量，不仅要声明，还要为他分配内存空间，否则值没办法存储
	// 值类型的声明不需要分配空间，因为在声明时已经默认分配好了内存空间
	// var j *int // 此处只声明，是个空指针，没有初始化，没有分配内存空间，之后的赋值操作没办法进行
	// *j = 100
	// fmt.Println(*j)

	//new函数申请一个内存地址
	// func new(Type) *Type
	// Type表示类型，new函数只接受一个参数，这个参数是个类型
	// *Type表示类型指针，new函数返回一个指向该类型内存地址的指针
	// new函数不太常用，给基本值类型指针申请内存，使用new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值
	var j1 *int
	fmt.Println(j1) // nil
	var j2 = new(int)
	fmt.Println(j2)  // 输出内存地址，此时就分配了内存
	fmt.Println(*j2) // 0 new返回的int指针对应的是int类型的零值
	*j2 = 101
	fmt.Println(*j2)
	j1 = new(int)
	fmt.Println(j1)
	fmt.Println(*j1)

	//make 只用于slice、map、chan的内存创建
	// 而且返回的是这三个类型本身，而不是他们类型的指针，因为这三个类型本身就是引用类型，没有必要返回他们的指针
	// func make(t Type, size ...IntegerType) Type
	// 输入的第一个参数是一个类型，而不是一个值
	var k map[string]int
	fmt.Println(k)
	fmt.Printf("%T\n", k)
	fmt.Println(k == nil)
	k = make(map[string]int)
	k["Kim"] = 100
	fmt.Println(k)

	// map 映射
	// 内部使用散列表 hash实现
	// 无序的基于key-value的数据结构
	// 引用类型，必须初始化才能使用
	// map[KeyType]ValueType
	// KeyType 表示键的类型
	// ValueType 表示键对应的值的类型
	// map类型的变量默认初始值为nil，需要使用make函数分配内存
	// make(map[KeyType]ValueType, [cap])
	// cap表示map的容量，不是必须的，但是应该在初始化的时候指定一个合适的值，避免程序运行期间动态扩容

	k["allen"] = 101
	k["taeyeon"] = 309
	fmt.Println(k["tayeon"]) // 返回0，不存在这个key
	fmt.Println(k["taeyeon"])

	value, ok := k["Kim"] // 返回的第二个参数是一个布尔值，true代表有这个key，false代表没有这个key，一般用ok接收返回的布尔值
	if !ok {
		fmt.Println("查无此key")
	} else {
		fmt.Println(value)
	}

	// for range遍历，遍历map时的元素顺序与添加键值对的顺序无关
	for key, value := range k { // 遍历key和value
		fmt.Println(key, value)
	}

	for key := range k { // 只遍历key
		fmt.Println(key)
	}

	for _, value := range k { // 只遍历value
		fmt.Println(value)
	}

	// delete函数删除键值对
	// delete(map, key)
	delete(k, "taeyeon")
	fmt.Println(k)
	delete(k, "taeyeon") // 函数没有返回值，删除不存在的元素，函数不进行操作，没有任何反应

	// 按照指定顺序遍历map
	// 把key取出来排序，排序后再遍历map
	// 先将map的所有key取出来存入切片，把切片排序，按照排序后的顺序遍历map

	rand.Seed(time.Now().UnixNano()) // 初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) // 生成stu开头的字符串
		value := rand.Intn(100)          // 生成0-99的随机整数
		scoreMap[key] = value
	}

	// 取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	// 对切片进行排序
	sort.Strings(keys)
	// 按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}

	// 元素为map类型的切片
	var mapSlice = make([]map[string]string, 3)
	fmt.Println(mapSlice)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")

	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "kim"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "beijing"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}

	// 值为切片类型的map
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value2, ok2 := sliceMap[key]
	if !ok2 {
		value2 = make([]string, 0, 3)
	}
	value2 = append(value2, "北京", "上海", "深圳")
	sliceMap[key] = value2
	fmt.Println(sliceMap)


	// 练习题 统计 "how do you do" 中各单词出现的次数
	// 把字符串按照空格切割得到切片
	string1 := "how do you do"
	string2 := strings.Split(string1, " ")
	// 遍历切片存储到一个map
	m1 := make(map[string]int, 10)
	for _, w := range string2 {
		// 如果map中不存在w这个key，那么出现次数=1
		// 如果map中存在w这个key，那么出现次数
		if _, ok := m1[w]; !ok {
			m1[w] = 1
		} else {
			m1[w]++
		}
	}
	// 累加各单词出现的次数
	for key, value := range m1 {
		fmt.Println(key, value)
	}

	// 练习题 回文判断
	// 字符串从左往右读和从右往左读一样
	string3 := "黄山落叶松叶落山黄"
	// 先把字符串中的字符拿出来放到一个rune切片中
	r := make([]rune, 0, len(string3))
	for _, c := range string3 {
		r = append(r, c)
	}
	fmt.Println(r)
	// 再把第一个和最后一个字符进行对比，依次对比相对称的字符
	for i := 0; i < len(r)/2; i++ { //只需要判断前一半即可
		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")

}
