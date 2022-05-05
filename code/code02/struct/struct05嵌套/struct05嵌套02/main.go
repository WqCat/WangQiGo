package main

import "fmt"

//嵌套结构体的字段冲突
type Address struct {
	Province   string
	City       string
	UpdateTime string
}

type Email struct {
	Addr       string
	UpdateTime string
}

type Person struct {
	Name    string
	Gender  string
	Age     int
	Address //嵌套另外一个结构体
	Email
}

func main() {
	p1 := Person{
		Name:   "弋痕夕",
		Gender: "男",
		Age:    26,
		Address: Address{
			Province:   "徐州",
			City:       "丰县",
			UpdateTime: "2020-02-02",
		},
		Email: Email{
			Addr:       "温庙",
			UpdateTime: "2022-03-02",
		},
	}
	fmt.Printf("%#v\n", p1)
	fmt.Println(p1.Address.UpdateTime)
	fmt.Println(p1.Email.UpdateTime)
}
