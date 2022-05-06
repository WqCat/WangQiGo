package main

import (
	"fmt"
	"strings"
)

//匿名函数和闭包
//闭包：1.函数可以作为返回值
//2.闭包=函数+外部函数引用

//定义一个函数，他的返回值是一个函数
func a(name string) func() {
	return func() {
		fmt.Println("hello", name)
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func calc2(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}
func main() {
	//匿名函数
	sayHello := func() {
		fmt.Println("匿名函数1")
	}
	sayHello()
	//立即执行
	func() {
		fmt.Println("匿名函数2")
	}()
	//闭包
	fmt.Println("闭包1")
	r := a("wangqi")
	r() //相当于执行了a的内部函数
	fmt.Println("闭包2")
	r1 := makeSuffixFunc(".txt")
	ret := r1("wq")
	fmt.Println(ret)
	fmt.Println("闭包3")
	x, y := calc2(100)
	ret1 := x(200)
	fmt.Println(ret1)
	ret2 := y(200)
	fmt.Println(ret2)
}
