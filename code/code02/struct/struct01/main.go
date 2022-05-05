package main

import "fmt"

//结构体
type person struct {
	name, city string
	age        int8
}

func main() {
	var p person
	p.name = "wq"
	p.city = "徐州"
	p.age = 18
	fmt.Printf("p=%#v\n", p)
	fmt.Println(p.city)
	//匿名结构体
	fmt.Println("匿名结构体")
	var user struct {
		name    string
		married bool
	}
	user.name = "we"
	user.married = false
	fmt.Printf("%#v\n", user)
	//结构体指针
	fmt.Println("结构体指针")
	var p2 = new(person)
	fmt.Printf("%T\n", p2)
	// (*p2).name = "wq"
	// (*p2).city = "徐州"
	// (*p2).age = 23
	p2.name = "wq"
	p2.city = "徐州"
	p2.age = 23
	fmt.Printf("%#v\n", p2)

	//取结构体地址进行实例化
	fmt.Println("实例化")
	p3 := &person{}
	fmt.Printf("%T\n", p3)
	fmt.Printf("%#v\n", p3)
	p3.name = "大司马"
	p3.city = "芜湖"
	p3.age = 35
	fmt.Printf("%#v\n", p3)
}
