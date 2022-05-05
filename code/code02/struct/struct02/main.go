package main

import "fmt"

//结构体初始化

type person struct {
	name, city string
	age        int8
}

func main() {
	//1.键值对初始化
	p1 := person{
		name: "wq",
		city: "徐州",
	}
	fmt.Printf("%#v\n", p1)

	p2 := &person{
		name: "侠岚",
		age:  20,
	}
	fmt.Printf("%#v\n", p2)
	//2.值得列表进行初始化
	p3 := person{
		"山鬼谣",
		"徐州",
		28,
	}
	fmt.Printf("%#v\n", p3)
}
