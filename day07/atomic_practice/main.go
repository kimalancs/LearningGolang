package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)
//原子操作
// 代码中的加锁操作因为涉及内核态的上下文切换会比较耗时、代价比较高
// 针对基本数据类型我们还可以使用原子操作来保证并发安全
// 因为原子操作是Go语言提供的方法它在用户态就可以完成，因此性能比加锁操作更好
// Go语言中原子操作由内置的标准库sync/atomic提供。

// atomic包
// 读取操作
// func LoadInt32(addr *int32) (val int32)
// func LoadInt64(addr *int64) (val int64)
// func LoadUint32(addr *uint32) (val uint32)
// func LoadUint64(addr *uint64) (val uint64)
// func LoadUintptr(addr *uintptr) (val uintptr)
// func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer)

// 写入操作
// func StoreInt32(addr *int32, val int32)
// func StoreInt64(addr *int64, val int64)
// func StoreUint32(addr *uint32, val uint32)
// func StoreUint64(addr *uint64, val uint64)
// func StoreUintptr(addr *uintptr, val uintptr)
// func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)

// 修改操作
// func AddInt32(addr *int32, delta int32) (new int32)
// func AddInt64(addr *int64, delta int64) (new int64)
// func AddUint32(addr *uint32, delta uint32) (new uint32)
// func AddUint64(addr *uint64, delta uint64) (new uint64)
// func AddUintptr(addr *uintptr, delta uintptr) (new uintptr)

// 交换操作
// func SwapInt32(addr *int32, new int32) (old int32)
// func SwapInt64(addr *int64, new int64) (old int64)
// func SwapUint32(addr *uint32, new uint32) (old uint32)
// func SwapUint64(addr *uint64, new uint64) (old uint64)
// func SwapUintptr(addr *uintptr, new uintptr) (old uintptr)
// func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer)

// 比较并交换操作
// func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)
// func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool)
// func CompareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool)
// func CompareAndSwapUint64(addr *uint64, old, new uint64) (swapped bool)
// func CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool)
// func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool)


// 比较下互斥锁和原子操作的性能
// atomic包提供了底层的原子级内存操作
// 对于同步算法的实现很有用。这些函数必须谨慎地保证正确使用。除了某些特殊的底层应用，使用通道或者sync包的函数/类型实现同步更好。




// Counter ...
type Counter interface {
	Inc()
	Load() int64
}

// 普通版

// CommonCounter ...
type CommonCounter struct {
	counter int64
}

// Inc ...
func (c CommonCounter) Inc() {
	c.counter++
}

// Load ...
func (c CommonCounter) Load() int64 {
	return c.counter
}

// 互斥锁版

// MutexCounter ...
type MutexCounter struct {
	counter int64
	lock    sync.Mutex
}

// Inc ...
func (m *MutexCounter) Inc() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.counter++
}

// Load ...
func (m *MutexCounter) Load() int64 {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.counter
}

// 原子操作版

// AtomicCounter ...
type AtomicCounter struct {
	counter int64
}

// Inc ...
func (a *AtomicCounter) Inc() {
	atomic.AddInt64(&a.counter, 1)
}

// Load ...
func (a *AtomicCounter) Load() int64 {
	return atomic.LoadInt64(&a.counter)
}

func test(c Counter) {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(c.Load(), end.Sub(start))
}

func main() {
	c1 := CommonCounter{} // 非并发安全
	test(c1)
	c2 := MutexCounter{} // 使用互斥锁实现并发安全
	test(&c2)
	c3 := AtomicCounter{} // 并发安全且比互斥锁效率更高
	test(&c3)
}