package main

import "fmt"

type student struct {
	id    int
	name  string
	class string
}

func newStudent(id int, name string, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}

//学员学生管理类
type studentMar struct {
	allStudents []*student
}

//newStudentMar 是studentMar的构造函数
func newStudentMar() *studentMar {
	return &studentMar{
		allStudents: make([]*student, 0, 100),
	}
}

//1.添加
func (s *studentMar) addStudent(newStu *student) {
	s.allStudents = append(s.allStudents, newStu)
}

//2.编辑
func (s *studentMar) modifyStudent(newStu *student) {
	for i, v := range s.allStudents {
		if newStu.id == v.id {
			s.allStudents[i] = newStu //学号相同时，什么切片索引，直接赋值
			return
		}
	}
	fmt.Printf("输入的学生找不到，系统没有%d的学生", newStu.id)
}

//3.展示学生
func (s *studentMar) showStudent() {
	for _, v := range s.allStudents {
		fmt.Printf("学号:%d 姓名:%s 班级:%s\n", v.id, v.name, v.class)
	}
}
