package main

import "fmt"

//为任意类型添加方法

//基于内置的基本类型造一个自己的类型
type myInt int

func (i myInt) sayHi() {
	fmt.Println("hi")
}

func main() {
	var m2 myInt
	fmt.Println(m2)
	m2 = 100
	m2.sayHi()
}
