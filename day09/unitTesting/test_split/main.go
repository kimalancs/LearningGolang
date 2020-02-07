package main

import (
	"fmt"

	"github.com/kimalancs/LearningGolang/day09/unitTesting1/split_strings"
)

func main() {
	ret := split_strings.Split("babcbef", "b")
	fmt.Printf("%#v\n", ret)
	ret = split_strings.Split("bbb", "b")
	fmt.Printf("%#v\n",ret)
	ret = split_strings.Split("acb","b")
	fmt.Printf("%#v\n",ret)
}
