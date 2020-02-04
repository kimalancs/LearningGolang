package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写互斥锁

// 互斥锁是完全互斥的，但是有很多实际的场景下是读多写少的
// 当我们并发的去读取一个资源不涉及资源修改的时候是没有必要加锁的，这种场景下使用读写锁是更好的一种选择
// 读写锁在Go语言中使用sync包中的RWMutex类型。

// 读写锁分为两种：读锁和写锁
// 当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待；
// 当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待

var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

func write() {
	// lock.Lock()   // 加互斥锁
	rwlock.Lock() // 加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设写操作耗时10毫秒
	rwlock.Unlock()                   // 解写锁
	// lock.Unlock()                     // 解互斥锁
	wg.Done()
}

func read() {
	// lock.Lock()                  // 加互斥锁
	rwlock.RLock()               // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwlock.RUnlock()             // 解读锁
	// lock.Unlock()                // 解互斥锁
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}

// 需要注意的是读写锁非常适合读多写少的场景，如果读和写的操作差别不大，读写锁的优势就发挥不出来。
// 读的goroutine获取的是读锁，后续的goroutine能读不能写，读并发，写串行
// 写的goroutine获取的是写锁，后续的goroutine不管是读还是写都要等待获取锁
// 读写锁，写者是排他性的，一个读写锁同时只能有一个写者或多个读者（与CPU数相关），但不能同时既有读者又有写者

// 互斥锁保证同一时刻只有一个goroutine可以访问资源，其他goroutine都要阻塞
// 互斥锁结构体很简单，其接口只有一个加锁和一个解锁的方法
// 当value为空时就是一个解锁的互斥锁，也就是其他人都可以来使用。并且当互斥锁第一次使用的时候就不能再被复制  
// 写大于读时，使用互斥锁  

