package main

import (
	"fmt"
	"time"
)

/*
time
时间的显示和测量用的函数
日历计算采用公历

time.Time，时间类型
time.Now() 获取当前的时间对象
在时间对象的基础上，使用方法获取年月日时分秒等信息


时间戳
自1970年1月1日（08:00:00GMT）至当前时间的总毫秒数
也被称为Unix时间戳 Unix Timestamp
基于时间对象可以通过Unix方法获取

time.Unix()函数可以将时间戳转为时间格式

时间间隔
time.Duration() 一个类型
一段时间间隔，可表示的最长时间段大约290年
代表两个时间点之间经过的时间，以纳秒为单位

时间间隔类型的常量如下
const (
	Nanosecond Duration = 1
	Microsecond 		= 1000 * Nanosecond
	Millisecond			= 1000 * Microsecond
	Second				= 1000 * Millisecond
	Minute				= 60 * Second
	Hour				= 60 * Minute
)
没有定义一天或超过一天的单元，避免夏时制的时区切换的混乱

时间操作
基于时间类型对象的方法来进行

Add
func (t Time) Add(d Duration) Time
时间点t经过时间间隔d之后是什么时间
某个时间点后多久是什么时间，如某个时间点后两个小时是几点
时间t + 时间间隔d
参数是Duration类型，返回Time类型

Sub
func (t Time) Sub(u Time) Duration
参数时间点u到时间点t，经过了多长的间隔，如早上八点到某个时间，经过了几个小时
两个时间之间的差值
返回t - u
u是t之前的时间点
参数是Time类型，返回Duration类型
如果结果超出了Duration可以表达的最大值或最小值，将返回最大值或最小值

要获取时间点t - d，可以用t.Add(-d)
某个时间点之前多久是什么时间，如某个时间点之前两个小时是几点
Add方法的参数是正数，得到某时间间隔之后的时间点
负数，某时间间隔之前的时间点

Equal
func (t Time) Equal(u Time) bool
判断两个时间是否相同，会考虑时区的影响，不同时区标准的时间也可以正确比较
本方法与t == u 不同，这种方法还会比较地点和时区信息

Before
func (t Time) Before(u Time) bool
如果t的时间点在u之前，返回真
u为参照点，t before u

After
func (t Time) After(u Time) bool
如果t的时间点在u之后，返回真
u为参照点，t afer u


定时器
time.Tick(时间间隔)设置定时器，定时器本质上是一个通道

时间格式化
Format
func (t Time) Format(layout string) string
格式化参考模板不是常见的 Y-M-D H:M:S
而是使用Go诞生时间2006年1月2号15时4分（快速记忆 2006-1-2-3-4）
Mon Jan 2 15:04:05 -0700 MST 2006

如果想格式化为12小时，需指定PM
*/
func timestampDemo() int64 {
	now := time.Now()
	timestamp1 := now.Unix()     // 基于时间对象获取时间戳
	timestamp2 := now.UnixNano() // 纳秒时间戳
	fmt.Printf("current timestamp1: %v\n", timestamp1)
	fmt.Printf("current timestamp2: %v\n", timestamp2)
	return timestamp1
}

func timestampDemo2(timestamp int64) {
	timeObj := time.Unix(timestamp, 0) // 将时间戳转为时间格式
	fmt.Println(timeObj)
}

func tickDemo(){
	ticker := time.Tick(time.Second) // 定义一个1秒间隔的定时器
	for i := range ticker {
		fmt.Println(i) // 每秒都会执行的任务
	}
}

func formatDemo(){
	now := time.Now()
	fmt.Println(now.Format("Mon Jan 2 15:04:05 -0700 MST 2006"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2016/01/02")) // 年份有问题
	fmt.Println(now.Format("3:04 PM 2016/01/02")) // 年份有问题
	fmt.Println(now.Format("2006/01/02"))
}

func main() {
	now := time.Now()
	fmt.Printf("current time: %v\n", now)
	// 2020-01-25 20:32:59.255373 +0800 CST m=+0.000157365

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	fmt.Println(now.Location()) //
	fmt.Println(now.Zone())     // 时区信息
	fmt.Println(now.Day())      // 那一月的第几天
	fmt.Println(now.Weekday())  // 周几
	fmt.Println(now.YearDay())  // 那一年的第几天
	fmt.Println(now.Date())     // 返回年月日三个值
	fmt.Println(now.Clock())    // 返回时分秒三个值
	fmt.Println(now.UTC())		// 对应的UTC时间

	timestamp := timestampDemo()
	timestampDemo2(timestamp)

	seconds := 12
	fmt.Println(int64(time.Second / time.Millisecond)) // 1000 将Duration类型值表示为某时间单元的个数，用除法
	fmt.Println(time.Duration(seconds) * time.Second)  // 12s 将整数个时间单元表示成Duration类型值，用乘法

	fmt.Println(now.Add(time.Hour * 2)) // 当前时间+两个小时

	// tickDemo()
	formatDemo()

	// 解析字符串格式的时间
	local, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	//按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2020/01/31 9:30:54", local)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 单变量声明时不能重复声明，多变量声明时可以，只要有一个新变量即可
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))

}
