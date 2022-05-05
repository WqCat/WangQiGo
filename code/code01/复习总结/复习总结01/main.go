package main

import "fmt"

func main() {
	var ages [30]int
	ages = [30]int{1, 2, 3, 4, 5, 6}
	fmt.Println(ages)

	var age02 = [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(age02)

	var age03 = [...]int{1: 100, 20: 30}
	fmt.Println(age03)

	//二维数组
	var a1 = [...][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(a1)

	//数组是值类型
	fmt.Println("数组是值类型")
	x := [3]int{1, 2, 3}
	f1(x)
	fmt.Println(x)

	y := x
	y[1] = 300 //只修改了y不影响x
	fmt.Println(x)
	fmt.Println(y)

	//切片
	fmt.Println("切片**********************")
	var s1 []int
	fmt.Println(s1)
	fmt.Println(s1 == nil) //没有分配内存=nil，没有初始化
	s1 = []int{1, 2, 3}
	fmt.Println(s1)
	//make 初始化 内存分配
	s2 := make([]bool, 2, 4)
	fmt.Println(s2) //[false,flase]
	s3 := make([]bool, 0, 4)
	fmt.Println(s3 == nil) //没有分配内存
	//切片不存值，类似地址
	fmt.Println("切片不存值和copy***")
	s4 := s1
	//var s6 []int//nil,无内存,拷贝不出
	var s6 = make([]int, 3, 3)
	copy(s6, s1)
	fmt.Println(s4) //[1 2 3]
	s4[1] = 200
	fmt.Println(s4) //[1 200 3]
	fmt.Println(s1) //[1 200 3]
	fmt.Println(s6) //[1 2 3]

	fmt.Println("扩容")
	var s5 []int //nil
	s5 = make([]int, 1, 2)
	s5[0] = 100
	fmt.Println(s5)
	s5 = append(s5, 1) //自动初始化，可以扩容
	fmt.Println(s5)

}

func f1(a [3]int) {
	a[1] = 23
}
