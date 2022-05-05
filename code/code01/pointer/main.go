package main

import "fmt"

//指针
func main() {
	a := 10 //int
	b := &a //*int
	fmt.Printf("%v %p\n", a, &a)
	fmt.Printf("%v %p %v\n", b, b, *b) //%v为值 %p地址  %T类型
	//

}
