package main

import "fmt"

//defer
//第一步:返回值赋值
//defer
//第二步：真真的RET返回

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++ //下面的类似当参数传进去，改变的是函数中的副本
	}(x)
	return 5
}

func f5() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

//传x的指针
func f6() (x int) {
	defer func(x *int) {
		(*x)++
	}(&x)
	return 5 //1.返回值x=5  2.defer x=6  3.RET返回
}

func main() {
	fmt.Println(f1()) //5
	fmt.Println(f2()) //6
	fmt.Println(f3()) //5
	fmt.Println(f4()) //5
	fmt.Println(f5()) //5
	fmt.Println(f6()) //6
}
