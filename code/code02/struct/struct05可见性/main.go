package main

import (
	"encoding/json"
	"fmt"
)

//结构体字段的可见性和JSON序列化
//定义的标识符首字母大写（Student和Name）,那么对外可见
//如果一个结构体的字段名首字母是大写的，那么该字符就是对外可见的
type student struct {
	ID   int
	Name string
}

func newStudent(id int, name string) student {
	return student{
		ID:   id,
		Name: name,
	}
}

type class struct {
	Title    string
	Teacher  string    `json:"teacher"`
	Students []student `json:"teacher_list" db:"student" xml:"stud_list"`
}

func main() {
	c1 := class{
		Title:    "高三七班",
		Teacher:  "王景梅",
		Students: make([]student, 0, 20),
	}
	//往c1添加学生
	for i := 0; i < 10; i++ {
		tmpStu := newStudent(i, fmt.Sprintf("stu%02d", i))
		c1.Students = append(c1.Students, tmpStu)
	}
	fmt.Printf("%#v\n", c1)
	//JSON序列化:GO语言中的数据-->JSON格式的字符串
	date, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json marshal failed,err", err)
		return
	}
	fmt.Println("序列化之后")
	fmt.Printf("%T\n", date)
	fmt.Printf("%s\n", date)
	//JSON反序列化：JSON格式的字符串--》GO语言中的数据
	fmt.Println("反序列化")
	jsonStr := `{"Title":"高三七班","Students":[{"ID":0,"Name":"stu00"},{"ID":1,"Name":"stu01"},{"ID":2,"Name":"stu02"},{"ID":3,"Name":"stu03"},{"ID":4,"Name":"stu04"}]}`
	var c2 class
	json.Unmarshal([]byte(jsonStr), &c2)
	if err != nil {
		fmt.Println("json unmarshal failed,err", err)
		return
	}
	fmt.Printf("%#v\n", c2)

}
