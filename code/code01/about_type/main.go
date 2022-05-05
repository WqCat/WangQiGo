package main

import "fmt"

//自定义类型和类型别名

//1.自定义类型

//MyInt 基于int类型的自定义类型
type MyInt int

//2.类型别名
//NewInt int的别名
type NewInt = int

func main() {
	var a MyInt
	fmt.Printf("type:%T value:%v\n", a, a)
	var b NewInt
	fmt.Printf("type:%T value:%v\n", b, b)
}
