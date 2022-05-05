package main

import "fmt"

//方法的实例化
type Person struct {
	name string
	age  int8
}

//Person类型的构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

//为Person定义方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是what", p.name)
}

//修改年龄的方法
//指针接受者(值会变)
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}

//值接受者来修改年龄（值不变）
func (p Person) SetAge2(newAge int8) {
	p.age = newAge
}

func main() {
	p1 := NewPerson("wq", int8(18))
	p1.Dream()

}
