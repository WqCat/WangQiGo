package main

import "fmt"

//闭包
func f(x, y int) {
	fmt.Println("this is f函数")
	fmt.Println(x + y)
}

func lixiang(x func(int, int), m, n int) func() {
	temp := func() {
		x(m, n)
	}
	return temp
}

func main() {
	ret := lixiang(f, 100, 200)
	ret()
}
