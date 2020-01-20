package smgr

import "fmt"

// Student student
type Student struct {
	ID   int64
	Name string
}

// StudentMgr student management
type StudentMgr struct {
	AllStudent map[int64]Student
}

// ShowMenu show the system menu
func ShowMenu() {
	fmt.Println("Welcome to SMS")
	fmt.Println(`
	1.查看所有学生信息
	2.添加学生信息
	3.修改学生信息
	4.删除学生信息
	5.退出
	`)
}

// ShowAllStudent show all student
func (s StudentMgr) ShowAllStudent() {
	// 从allStudent中遍历所有学生信息
	for _, stu := range s.AllStudent {
		fmt.Printf("学号：%d 姓名：%s\n", stu.ID, stu.Name)
	}
}

// AddStudent add a new student
func (s StudentMgr) AddStudent() {
	fmt.Println(s)
	// 根据用户输入的内容，创建一个新的学生
	var (
		stuID   int64
		stuName string
	)
	// 获取用户输入
	fmt.Println("请输入学号：")
	fmt.Scanln(&stuID)
	fmt.Println("请输入姓名：")
	fmt.Scanln(&stuName)
	newStu := Student{
		ID:   stuID,
		Name: stuName,
	}
	// 将新学生信息添加到s.AllStudent这个map中
	s.AllStudent[newStu.ID] = newStu
	fmt.Println("添加成功")
}

// EditStudent Edit the information of a student
func (s StudentMgr) EditStudent() {
	// 获取用户输入的学号
	var stuID int64
	fmt.Println("请输入要修改学生信息的学号：")
	fmt.Scanln(&stuID)
	//展示该学号对应的学生信息
	stuObj, ok := s.AllStudent[stuID]
	if !ok {
		fmt.Println("没有找到该学号对应的学生信息")
		return
	}
	fmt.Printf("你要修改的学生信息如下：学号：%d 姓名：%s\n", stuObj.ID, stuObj.Name)
	fmt.Println("请输入学生的新名字：")
	// 获取修改后的学生名
	fmt.Scanln(&stuObj.Name)
	// 更新学生名
	s.AllStudent[stuID] = stuObj
}

//DeleteStudent delete a student
func (s StudentMgr) DeleteStudent() {
	// 获取需删除学生信息的学号
	var stuID int64
	fmt.Println("请输入要删除学生的学号：")
	fmt.Scanln(&stuID)
	//展示该学号对应的学生信息
	stuObj, ok := s.AllStudent[stuID]
	if !ok {
		fmt.Println("没有找到该学号对应的学生信息")
		return
	}
	fmt.Printf("你要删除的学生信息如下：学号：%d 姓名：%s\n", stuObj.ID, stuObj.Name)
	fmt.Println(`
	是否删除
	1.确认删除
	2.返回菜单
	`)
	var confirm int
	fmt.Scanln(&confirm)
	switch confirm {
	case 1:
		delete(s.AllStudent, stuObj.ID)
	case 2:
		fmt.Println("返回上级菜单")
	}
}
