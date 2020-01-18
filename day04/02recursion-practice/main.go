package main

import "fmt"

// 上台阶
// n个台阶，一次可以走1步，也可以走2步，有多少种走法

func taijie(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return taijie(n-1) + taijie(n-2)
}
func main() {
	fmt.Println(taijie(10))

}
