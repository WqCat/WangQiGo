package main

import "fmt"

//使用值接受者实现接口和使用指针接受者实现接口的区别

//接口的嵌套
type animal interface {
	mover
	sayer
}

type mover interface {
	move()
}

type sayer interface {
	say()
}

type person struct {
	name string
	age  int8
}

//使用值接受者实现接口:类型的值和类型的指针都能保存在接口变量中
// func (p person) move() {
// 	fmt.Printf("%s再跑\n", p.name)
// }

// 使用指针接受者实现接口
func (p *person) move() {
	fmt.Printf("%s再跑\n", p.name)
}

func (p *person) say() {
	fmt.Printf("%s在狗叫什么\n", p.name)
}

func main() {
	var m2 mover
	// p1 := person{ //person类型的值
	// 	name: "盖聂",
	// 	age:  22,
	// }
	p2 := &person{ //person类型的指针
		name: "卫庄",
		age:  20,
	}
	//m1 = p1 //指针接受的，无法赋值，因为person类型的值没有实现mover的接口
	m2 = p2
	// fmt.Println("值接受者--------------")
	// m1.move()
	// fmt.Println(m1)
	// m2.move()
	// fmt.Println(m2)

	fmt.Println("指针接受者--------------")
	// m1.move()
	// fmt.Println(m1)
	m2.move()
	fmt.Println(m2)

	var s sayer
	s = p2
	s.say()
	fmt.Println(s)

	fmt.Println("接口嵌套-------------")
	var a animal
	a = p2
	a.move()
	a.say()
	fmt.Println(a)
}
