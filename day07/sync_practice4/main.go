package main

import (
	"fmt"
	"strconv"
	"sync"
)

/*
在编程的很多场景下我们需要确保某些操作在高并发的场景下只执行一次，例如只加载一次配置文件、只关闭一次通道等

Go语言中的sync包中提供了一个针对只执行一次场景的解决方案:sync.Once

sync.Once只有一个Do方法，其签名如下：
func (o *Once) Do(f func()) {}
Do()传入的参数只能是没有参数也没有返回值的函数
如果要执行的函数f需要传递参数就需要搭配闭包来使用
使用匿名函数包装一个有参数的函数赋值给一个变量，然后把这个函数变量传给Do()方法


*/

// 加载配置文件示例
// 延迟一个开销很大的初始化操作到真正用到它的时候再执行是一个很好的实践
// 因为预先初始化一个变量（比如在init函数中完成初始化）会增加程序的启动耗时
// 而且有可能实际执行过程中这个变量没有用上，那么这个初始化操作就不是必须要做的。我们来看一个例子：

// var icons map[string]image.Image

// func loadIcons() {
// 	icons = map[string]image.Image{
// 		"left":  loadIcon("left.png"),
// 		"up":    loadIcon("up.png"),
// 		"right": loadIcon("right.png"),
// 		"down":  loadIcon("down.png"),
// 	}
// }

// Icon 被多个goroutine调用时不是并发安全的
// func loadIcon(name string) image.Image {
// 	if icons == nil {
// 		loadIcons()
// 	}
// 	return icons[name]
// }

// 多个goroutine并发调用Icon函数时不是并发安全的
// 现代的编译器和CPU可能会在保证每个goroutine都满足串行一致的基础上自由地重排访问内存的顺序
// loadIcons函数可能会被重排为以下结果：

// func loadIcons() {
// 	icons = make(map[string]image.Image)
// 	icons["left"] = loadIcon("left.png")
// 	icons["up"] = loadIcon("up.png")
// 	icons["right"] = loadIcon("right.png")
// 	icons["down"] = loadIcon("down.png")
// }

// 在这种情况下就会出现即使判断了icons不是nil也不意味着变量初始化完成了(make初始化后icons就不是nil了)
// 考虑到这种情况，我们能想到的办法就是添加互斥锁，保证初始化icons的时候不会被其他的goroutine操作，但是这样做又会引发性能问题。

// 用sync.Once改造的示例代码如下：

// var icons map[string]image.Image

// var loadIconsOnce sync.Once

// func loadIcons() {
// 	icons = map[string]image.Image{
// 		"left":  loadIcon("left.png"),
// 		"up":    loadIcon("up.png"),
// 		"right": loadIcon("right.png"),
// 		"down":  loadIcon("down.png"),
// 	}
// }

// Icon 是并发安全的
// func Icon(name string) image.Image {
// 	loadIconsOnce.Do(loadIcons) // Do()方法只能接收没有参数没有接收值的函数
// 	return icons[name]
// }

// 下面是借助sync.Once实现的并发安全的单例模式：

// package singleton

// import (
//     "sync"
// )

// type singleton struct {}

// var instance *singleton
// var once sync.Once

// func GetInstance() *singleton {
//     once.Do(func() {
//         instance = &singleton{}
//     })
//     return instance
// }

// sync.Once其实内部包含一个互斥锁和一个布尔值
// 互斥锁保证布尔值和数据的安全，而布尔值用来记录初始化是否完成
// 这样设计就能保证初始化操作的时候是并发安全的并且初始化操作也不会被执行多次。

// sync.Map
// Go语言中内置的map不是并发安全的。请看下面的示例：

// var m = make(map[string]int)

// func get(key string) int {
// 	return m[key]
// }

// func set(key string, value int) {
// 	m[key] = value
// }

// func main() {
// 	wg := sync.WaitGroup{}
// 	for i := 0; i < 20; i++ { // 20个goroutine并发就是map的极限，超过20就会报错
// 		wg.Add(1)
// 		go func(n int) {
// 			key := strconv.Itoa(n)
// 			set(key, n)
// 			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
// 			wg.Done()
// 		}(i)
// 	}
// 	wg.Wait()
// }
// 上面的代码开启少量几个goroutine的时候可能没什么问题
// 当并发多了之后（超过20个goroutine）执行上面的代码就会报fatal error: concurrent map writes错误

// 像这种场景下就需要为map加锁来保证并发的安全性了
// Go语言的sync包中提供了一个开箱即用的并发安全版map：sync.Map
// 开箱即用表示不用像内置的map一样使用make函数初始化就能直接使用
// 同时sync.Map内置了诸如Store、Load、LoadOrStore、Delete、Range等操作方法

var m = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
