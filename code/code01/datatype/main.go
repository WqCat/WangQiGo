package main

import (
	"fmt"
	"math"
)

//基本数据类型
func main() {
	//十进制打印二进制
	n := 10
	fmt.Printf("%b\n", n)
	fmt.Printf("%d\n", n)
	//八进制
	m := 075
	fmt.Printf("%o\n", m)
	//十六进制
	f := 0xff
	fmt.Println(f)
	fmt.Printf("%x\n", f)

	//uint8
	var age uint8 //0~255
	age = 255
	fmt.Println(age)

	//浮点数
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	//字符串
	s1 := "hello bejing"
	s2 := "你好 徐州"
	fmt.Println(s1)
	fmt.Println(s2)

	//打印路径
	fmt.Println("c:\\code\\love\\you\\who")
	fmt.Println("\t制表符\n换行符")
	s3 := `
	多行字符串
	两个反引号
	回原样输出
	`
	fmt.Println(s3)
}
