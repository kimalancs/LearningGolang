package main

import (
	"fmt"
	"os"
	"github.com/kimalancs/LearningGolang/day04/06practice2/smgr"
)


func main() {
	var smgr1 = smgr.StudentMgr{
		AllStudent: make(map[int64]smgr.Student, 100),
	}
	for {
		smgr.ShowMenu()
		// 等待用户输入选项
		fmt.Println("请输入序号：")
		var choice int
		fmt.Scanln(&choice)
		fmt.Println("你输入的是：", choice)

		switch choice {
		case 1:
			smgr1.ShowAllStudent()
		case 2:
			smgr1.AddStudent()
		case 3:
			smgr1.EditStudent()
		case 4:
			smgr1.DeleteStudent()
		case 5:
			os.Exit(0)
		default:
			fmt.Println("选择错误")
		}
	}
}
