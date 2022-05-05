package main

import (
	"fmt"
)

//空接口
//接口没有定义任何方法
//任意类型都实现了空接口————》空接口可以存储任意值

//空接口一般不需要提前定义
func main() {
	var x interface{} //定义空接口
	x = "hello"
	fmt.Println(x)
	x = false
	fmt.Println(x)

	var m = make(map[string]interface{}, 16)
	m["name"] = "天明"
	m["age"] = 18
	m["hobby"] = []string{"乒乓球", "羽毛球"}
	fmt.Println(m)

	//类型断言
	ret, ok := x.(string) //类型断言,猜的不对就返回一个布尔型  ok=false,ret=string类型的零值
	if !ok {
		fmt.Println("不是字符串类型")
	} else {
		fmt.Println("不是字符串类型")
	}
	fmt.Println(ret)

	//使用switch进行类型断言
	switch v := x.(type) {
	case string:
		fmt.Printf("是字符串类型,value:%v\n", v)
	case bool:
		fmt.Printf("是布尔类型,value:%v\n", v)
	case int:
		fmt.Printf("是int类型,value:%v\n", v)
	default:
		fmt.Printf("猜不到,value:%v\n", v)
	}

}
