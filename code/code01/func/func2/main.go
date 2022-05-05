package main

import "fmt"

//函数进阶值变量作用域

var num = 10

func testGlobal() {
	num := 100
	fmt.Println("全局变量", num) //函数中有的话，不找去全局变量了
}
func main() {
	fmt.Println(num)
	testGlobal()
	abc := testGlobal
	fmt.Printf("%T\n", abc)
	abc()
	fmt.Println("复合")
	r1 := calc(100, 200, add)
	fmt.Println(r1)
}
func add(x, y int) int {
	return x + y
}

func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}
