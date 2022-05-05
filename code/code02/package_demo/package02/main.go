package main

import (
	wq "code/code02/package_demo/package01" //给导入的包起别名
	"fmt"
)

//Go语言不循序导入包却不使用
//GO语言不允许循环使用

// import "code/code02/package_demo/package01"

//多用来做一些初始化操作
func init() {
	fmt.Println("main.init()")
}

func main() {
	fmt.Println("hello")
	ret := wq.Add(10, 20)
	fmt.Println(ret)
	fmt.Println(wq.Name)
}
