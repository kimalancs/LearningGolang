package main

import (
	"fmt"
	"os"
)

/*
	student management system
	结构体方法版学生管理系统
	查看、新增、删除学生信息

*/

type student struct {
	id   int64
	name string
}

type studentMgr struct {
	allStudent map[int64]student
}

func showMenu() {
	fmt.Println(`
	Welcome to SMS 
	1.查看所有学生信息
	2.添加学生信息
	3.修改学生信息
	4.删除学生信息
	5.退出
	`)
}

func (s studentMgr) showAllStudent() {
	// 从allStudent中遍历所有学生信息
	for _, stu := range s.allStudent {
		fmt.Printf("学号：%d 姓名：%s\n", stu.id, stu.name)
	}
}

func (s studentMgr) addStudent() {
	// 根据用户输入的内容，创建一个新的学生
	var (
		stuID   int64
		stuName string
	)
	// 获取用户输入
	fmt.Println("请输入学号：")
	fmt.Scanln(&stuID)
	for _, v := range s.allStudent {
		if stuID == v.id {
			fmt.Println("该学号已存在")
			return
		}
	}
	fmt.Println("请输入姓名：")
	fmt.Scanln(&stuName)
	newStu := student{
		id:   stuID,
		name: stuName,
	}
	// 将新学生信息添加到s.AllStudent这个map中
	s.allStudent[newStu.id] = newStu
	fmt.Println("添加成功")
}

func (s studentMgr) editStudent() {
	// 获取用户输入的学号
	var stuID int64
	fmt.Println("请输入要修改学生信息的学号：")
	fmt.Scanln(&stuID)
	//展示该学号对应的学生信息
	stuObj, ok := s.allStudent[stuID]
	if !ok {
		fmt.Println("没有找到该学号对应的学生信息")
		return
	}
	fmt.Printf("你要修改的学生信息如下：学号：%d 姓名：%s\n", stuObj.id, stuObj.name)
	fmt.Println("请输入学生的新名字：")
	// 获取修改后的学生名
	fmt.Scanln(&stuObj.name)
	// 更新学生名
	s.allStudent[stuID] = stuObj
}

func (s studentMgr) deleteStudent() {
	// 获取需删除学生信息的学号
	var stuID int64
	fmt.Println("请输入要删除学生的学号：")
	fmt.Scanln(&stuID)
	//展示该学号对应的学生信息
	stuObj, ok := s.allStudent[stuID]
	if !ok {
		fmt.Println("没有找到该学号对应的学生信息")
		return
	}
	fmt.Printf("你要删除的学生信息如下：学号：%d 姓名：%s\n", stuObj.id, stuObj.name)
	fmt.Println(`
	确认是否删除该学生信息，请输入序号：
	1.确认删除
	2.返回菜单
	`)

	var confirm int
	fmt.Scanln(&confirm)
	switch confirm {
	case 1:
		delete(s.allStudent, stuObj.id)
	case 2:
		fmt.Println("返回上级菜单")
	}
}

func main() {
	var smgr = studentMgr{
		allStudent: make(map[int64]student, 100),
	}
	for {
		showMenu()
		// 等待用户输入选项
		fmt.Println("请输入序号：")
		var choice int
		fmt.Scanln(&choice)
		fmt.Println("你输入的是：", choice)

		switch choice {
		case 1:
			smgr.showAllStudent()
		case 2:
			smgr.addStudent()
		case 3:
			smgr.editStudent()
		case 4:
			smgr.deleteStudent()
		case 5:
			os.Exit(0)
		default:
			fmt.Println("选择错误")
		}
	}
}
