package main

import (
	"fmt"
	"os"
)

/*
	student management system
	函数版学生管理系统
	查看、新增、删除学生信息

*/

type student struct {
	id    int64
	name  string
	age   int64
	score float64
}

var (
	allStudent map[int64]*student
)

// 打印所有学生信息
func showAllStudent() {
	for _, v := range allStudent {
		fmt.Printf("学号：%d 姓名 %s\n", v.id, v.name)
	}

}

// 学生信息，结构体构造函数
func newStudent(id int64, name string) *student {
	return &student {
		id: id, 
		name: name,
	}
}

// add a new student to the list
func addStudent() {
	// 创建一个学生信息
	var (
		id int64
		name string
	)
	fmt.Println("请输入学生学号：")
	fmt.Scanln(&id)
	fmt.Println("请输入学生姓名")
	fmt.Scanln(&name)
	// 调用构造函数
	newStu := newStudent(id, name)
	// 追加到allStudent
	allStudent[id] = newStu
	fmt.Printf("%s学生信息已添加\n", newStu.name)
}

func deleteStudent() {
	// 用户输入要删除的学生的学号
	var (
		deletedID int64
	)
	fmt.Print("请输入要删除的学生学号：")
	fmt.Scanln(&deletedID)
	// 删除allStudent中对应的键值对
	delete(allStudent, deletedID)

}
// // show the infomation of the student
// func (s *student) info() {
// 	fmt.Printf("学生基本信息：Id %04d  姓名 %s 年龄 %d岁\n", s.id, s.name, s.age)
// }

// func (s *student) changeID(id int) {
// 	s.id = id
// }

// func (s *student) changeName(name string) {
// 	s.name = name
// }

// func changeInfo(stu *student) {
// 	fmt.Printf("学生基本信息：Id %04d  姓名 %s 年龄 %d岁\n", stu.id, stu.name, stu.age)
// 	fmt.Println("请输入学生id：")
// 	_, err := fmt.Scanf("%d\n", &stu.id)
// 	if err != nil {
// 		fmt.Println("请输入数字格式")
// 		changeInfo(stu)
// 	}
// 	fmt.Println("请输入学生姓名：")
// 	fmt.Scanln(&stu.name)
// 	fmt.Println("请输入新学生年龄：")
// 	fmt.Scanln(&stu.age)

// 	fmt.Printf("学生基本信息更新为：Id %04d  姓名 %s 年龄 %d岁\n", stu.id, stu.name, stu.age)

// }

func main() {
	allStudent = make(map[int64]*student, 100) // 初始化，开辟内存空间
	for {
		// 打印菜单
		fmt.Println("欢迎使用学生管理系统")
		fmt.Println(`
		1.查看所有学生信息
		2.新增学生信息
		3.删除学生信息
		4.退出管理系统
    	`)
		fmt.Print("请选择你要执行的操作：")
		// 用户选择操作
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了操作%d\n", choice)
		// 执行对应操作
		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(0) //退出
		default:
			fmt.Println("错误选项")
		}
	}

}
