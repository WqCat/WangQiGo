package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//map(映射)
func main() {
	//光声明map
	var a map[string]int
	fmt.Println(a == nil)
	//map初始化
	a = make(map[string]int, 8)
	fmt.Println(a == nil)
	//map添加键值对
	a["徐州"] = 100
	a["南京"] = 200
	fmt.Println(a)
	fmt.Printf("a:%v\n", a)
	fmt.Printf("a:%#v\n", a)
	fmt.Printf("type:%T\n", a)
	//声明并且初始化
	b := map[int]bool{
		1: true,
		2: false,
	}
	fmt.Printf("b:%#v\n", b)
	fmt.Printf("type:%T\n", b)

	// var c map[int]int
	// c[100]=200     //错误，c没有初始化，不能直接操作
	// fmt.Println(c)

	//判断键存不存在
	var scoreMap = make(map[string]int, 8)
	scoreMap["徐州"] = 100
	scoreMap["南京"] = 200

	value, ok := scoreMap["徐州"]
	fmt.Println(value, ok)
	if ok {
		fmt.Println("徐州在", value)
	} else {
		fmt.Println("上海不在")
	}

	//map遍历 for range
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
	for k := range scoreMap {
		fmt.Println(k)
	}

	//删除 南京
	delete(scoreMap, "南京")
	fmt.Println(scoreMap)

	//按照固定顺序遍历map
	var cityMap = make(map[string]int, 100)

	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100) //0-99随机整数
		cityMap[key] = value
	}
	for k, v := range cityMap {
		fmt.Println(k, v)
	}
	//按照key从小到大遍历
	keys := make([]string, 0, 100)
	for k := range cityMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, cityMap[key])
	}
}
