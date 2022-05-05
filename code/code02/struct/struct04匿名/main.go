package main

import "fmt"

//结构体匿名字段
type Person struct {
	string
	int
}

func main() {
	p1 := Person{
		"辗迟",
		18,
	}
	fmt.Println(p1)
	fmt.Println(p1.string, p1.int)
}
