package main

import "fmt"

//为什么用接口

type dog struct {
}

func (d dog) say() {
	fmt.Println("汪汪汪。。")
}

type cat struct {
}

func (c cat) say() {
	fmt.Println("喵喵喵。。")
}

type person struct {
	name string
}

func (p person) say() {
	fmt.Println("啊啊啊啊啊")
}

//interface不管你是什么类型，只管实现什么方法
//定义一个类型，只要能实现sau()都可以成为sayer类型
type sayer interface {
	say()
}

//打的函数
func da(arg sayer) {
	arg.say() //传进来都执行say
}

func main() {
	c1 := cat{}
	da(c1)
	d1 := dog{}
	da(d1)
	p1 := person{
		name: "假叶",
	}
	da(p1)

	var s sayer
	c2 := cat{}
	s = c2
	p2 := person{
		name: "王一帆",
	}
	s = p2
	fmt.Println(s)
}
