package main

import (
	"fmt"

	"github.com/kimalancs/LearningGolang/day09/unitTesting/splitstrings"
)

func main() {
	ret := splitstrings.Split("babcbef", "b")
	fmt.Printf("%#v\n", ret)
	ret = splitstrings.Split("bbb", "b")
	fmt.Printf("%#v\n", ret)
	ret = splitstrings.Split("acb", "b")
	fmt.Printf("%#v\n", ret)
}
