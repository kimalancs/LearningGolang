package main

import ( // 多行导入格式
	"fmt"

	s "github.com/studygo/day04/06practice2/smgr" // 自定义包名，当包名太长或冲突时使用

)
/*
package
包
组织代码的单位
高级的代码复用方案

main包可以编译生成可执行文件，其他包不可以

一个包就是一个存放.go文件的文件夹
该文件夹下所有的go文件都要在代码第一行声明该文件归属的包
package 包名
一个包的代码可以分成多个.go文件
编译时会一起编译

一个文件夹下面包含的文件只能归属于一个package
只在本级目录下，如果有子目录，就是另一个包了

同一个package的文件不能在多个文件夹下

包名可以和文件夹名字不一样，包名不能包含 - 符号（横线）
包名和标识符命名一样，不能以数字开头

标识符
首字母大写，对外可见的、public、其他包可以访问
首字母小写，外部包不可见，只能在当前包内使用

函数内部的局部变量，首字母大写，也只能在当前函数内使用

结构体中的字段名和接口中的方法名，首字母大写，外部包可访问

import "包的路径"
导入包时写的是包的路径，最后一部分是包文件夹名，而不是包名
import "github.com/kimalancs/studygo/day04/06practice2/smgr"

import语句放在文件开头包声明语句下边
双引号包裹
导入包路径，从¥GOPATH/src/后开始计算
禁止循环导入包

导入的包未使用无法编译

匿名导入包
只希望导入包，而不使用包内部数据时使用
匿名导入的包和其他包一样会编译到可执行文件中

init函数
Go语言执行时导入包语句会自动触发包内部init()函数的调用
init()函数没有参数也没有返回值
init()函数在程序运行时自动调用执行，不能在代码中主动调用它

全局声明 ==> init() ==> main()

从main包开始检查其导入的包，每个包又可能导入了其他包
构建一个树状的包引用关系，根据引用顺序决定编译顺序
最后导入的包最先初始化并调用其init()函数

main ==> A ==> B ==> C
main包导入A，A导入B，B导入C

C.init() ==> B.init() ==> A.init() ==> main.init()

*/

func main() {
	var stu s.Student 
	fmt.Printf("%T\n",stu)

	stu = s.Student{ID:18,Name:"Taeyeon"}
	fmt.Println(stu)
}