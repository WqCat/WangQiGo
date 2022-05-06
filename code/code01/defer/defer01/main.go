package main

import "fmt"

//defer:延迟执行
func main() {
	fmt.Println("start...")
	defer fmt.Println(1) //类似栈结构，先进后出
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end..")
}
