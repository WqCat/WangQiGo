package main

import "fmt"

var xy = 100

//变量
func main() {
	//标准声明格式
	var name string
	var age int
	var isOk bool
	fmt.Println(name, age, isOk)
	//批量声明
	var (
		a string
		b int
		c bool
		d float32
	)
	fmt.Println(a, b, c, d)
	//设置初始值
	var name1 string = "花木兰"
	var age1 int = 18
	fmt.Println(name1, age1)
	//设置初始值简写
	var name3 = "元歌"
	var age3 = 23
	fmt.Println(name3, age3)
	//短变量声明
	m := 10
	fmt.Println(m)
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)
}
func foo() (int, string) {
	return 10, "W1Q1"
}
