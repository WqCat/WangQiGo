package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y)) // 先把里面的calc()执行完输入，然后接着往下执行，最后最外层的calc()
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}

//1.defer calc("AA",1,calc("A",1,2))
//2.calc("A",1,2)   ****************--->> A 1 2 3
//3.defer calc("AA",1,3)
//4.defer calc("BB", 10, calc("B", 10,2))
//5.calc("B", 10,2)  ********--->> B 10 2 12
//6.defer calc("BB", 10, 12)
//7.执行defer calc("BB", 10, 12) *********--->> BB 10 12 22
//8.执行defer calc("AA",1,3) ******--->> AA 1 3 4
