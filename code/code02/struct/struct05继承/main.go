package main

import "fmt"

//结构体继承

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动\n", a.name)
}

type Dog struct {
	Feet    int8
	*Animal //匿名嵌套指针
}

func (d *Dog) wang() {
	fmt.Printf("%s会旺旺叫\n", d.name)
}

func main() {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{
			name: "金毛",
		},
	}
	d1.wang()
	d1.move()
}
