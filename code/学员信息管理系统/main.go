package main

import (
	"fmt"
	"os"
)

//学员信息管理系统、

//1.学员信息增删改查

func showMenu() {
	fmt.Println("欢迎来到学员信息管理系统")
	fmt.Println("1.添加学员")
	fmt.Println("2.编辑学员")
	fmt.Println("3.展示所有学员信息")
	fmt.Println("4.退出系统")
}

//获取学员信息
func getInput() *student {
	var (
		id    int
		name  string
		class string
	)
	fmt.Println("请输入学员信息：")
	fmt.Printf("请输入学员学号:")
	fmt.Scanf("%d\n", &id)
	fmt.Printf("请输入学员姓名:")
	fmt.Scanf("%s\n", &name)
	fmt.Printf("请输入学员班级:")
	fmt.Scanf("%s\n", &class)
	stu := newStudent(id, name, class)
	return stu
}

func main() {
	sm := newStudentMar()
	for {
		showMenu()
		var input int
		fmt.Println("请输入想要操作的序号：")
		fmt.Scanf("%d\n", &input)
		fmt.Println("用户输入的为：", input)
		//执行动作
		switch input {
		case 1:
			//添加学员
			stu := getInput()
			sm.addStudent(stu)
		case 2:
			//编辑学员
			stu := getInput()
			sm.modifyStudent(stu)
		case 3:
			//展示所有学员
			sm.showStudent()
		case 4:
			//退出
			os.Exit(0)
		}
	}
}
