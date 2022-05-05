package main

import "fmt"

func main() {
	//指针
	//Go里面的指针只读不能修改，不能修改指针变量指向的地址
	addr := "徐州"
	addrP := &addr
	fmt.Println(addrP)
	fmt.Printf("%T\n", addrP)
	addV := addrP
	fmt.Println(addV)

	//map,需要声明内存
	var m1 map[string]int
	fmt.Println(m1 == nil)
	m1 = make(map[string]int, 10)
	m1["徐州"] = 100
	fmt.Println(m1)
	fmt.Println(m1["wq"]) //如果key不存在，返回对应类型的默认值
	//如果值是bool，通常用ok接收它
	score, ok := m1["ji"]
	if !ok {
		fmt.Println("没有此人")
	} else {
		fmt.Println("姬无命的分数是", score)
	}
	delete(m1, "we") //删除的key不存在什么都不干
	delete(m1, "徐州")
	fmt.Println(m1)
	fmt.Println(m1 == nil) //已经开辟了空间

}
