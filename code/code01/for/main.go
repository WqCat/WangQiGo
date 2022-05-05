package main

import "fmt"

//for循环
func main() {
	//1 基础版本
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 3 {
			break
		}
	}
	//
}
