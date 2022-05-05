package main

import "fmt"

//嵌套结构体

type Address struct {
	Province string
	City     string
}

type Person struct {
	Name    string
	Gender  string
	Age     int
	Address //嵌套另外一个结构体
}

func main() {
	p1 := Person{
		Name:   "山鬼谣",
		Gender: "男",
		Age:    28,
		Address: Address{
			Province: "徐州",
			City:     "丰县",
		},
	}
	fmt.Printf("%#v\n", p1)
	fmt.Println(p1.Name, p1.Gender, p1.Age, p1.Address)
	fmt.Println(p1.Address.Province)
	fmt.Println(p1.Province)
}
