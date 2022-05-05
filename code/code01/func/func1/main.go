package main

import "fmt"

//函数

func sayHello() {
	fmt.Println("hello wq")
}

func sayHello2(name string) {
	fmt.Println("hello", name)
}

func intSum(a int, b int) int {
	ret := a + b
	return ret
}

func intSum2(a int, b int) (ret int) {
	ret = a + b
	return
}

//函数接收可变参数, ...表示可变参数
//可变参数是切片类型
func intSum3(a ...int) int {
	ret := 0
	for _, arg := range a {
		ret = ret + arg
	}
	return ret
}

//可变参数在固定参数后面
//go语言函数无默认参数
func intSum4(a int, b ...int) int {
	ret := a
	for _, arg := range b {
		ret = ret + arg
	}
	return ret
}

//go语言函数参数类型简写
func intSum5(a, b int) (ret int) {
	ret = a + b
	return
}

//定义多个返回值的函数,支持简写
func calc(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}
func main() {
	sayHello()
	name := "徐州"
	sayHello2(name)
	r := intSum(10, 20)
	fmt.Println(r)
	intSum3(10, 20)
	r1 := intSum3()
	r2 := intSum3(10, 20, 30)
	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println("多参数")
	r3 := intSum4(10)
	r4 := intSum4(10, 20, 30, 40) //a=10,b=[20,30,40]
	fmt.Println(r3)
	fmt.Println(r4)
	fmt.Println("多返回值")
	x, y := calc(100, 200)
	fmt.Println(x, y)

}
