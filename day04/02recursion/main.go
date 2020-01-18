package main

import "fmt"

// recursion 递归
// 自己调用自己
// 适合处理那种问题相同、问题规模越来越小的场景
// 递归一定要有一个明确的退出条件
// 阶乘
// 5！= 5*4*3*2*1

func f1(n uint64) uint64 {
	if n <= 0 {
		fmt.Println("请输入大于0的整数")
	} else if n == 1 {
		return 1
	}
	return n * f1(n-1)
}

func main() {
	fmt.Println(f1(20)) // 21以上出问题，结果为负数
	// var inputNumber uint64
	// fmt.Println("请输入一个正整数")
	// fmt.Scanf("%d", &inputNumber)
	// fmt.Printf("%d的阶乘结果为：%d\n", inputNumber, f1(inputNumber))
}
