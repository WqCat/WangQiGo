package main

import (
	"fmt"
	"strings"
)

//map
func main() {
	//元素为map的切片,
	var mapSlice = make([]map[string]int, 8, 20)
	mapSlice[0] = make(map[string]int, 8)
	mapSlice[0]["徐州"] = 100
	fmt.Println(mapSlice)

	//值为切片的Map
	var silceMap = make(map[string][]int, 8) //只完成了map的初始化
	v, ok := silceMap["徐州"]
	if ok {
		fmt.Println(v)
	} else {
		silceMap["徐州"] = make([]int, 8)
		silceMap["徐州"][0] = 100
		silceMap["徐州"][1] = 200
		silceMap["徐州"][3] = 300
	}
	for k, v := range silceMap {
		fmt.Println(k, v)
	}
	fmt.Println("练习题1")
	//统计字符串单词出现的次数
	//每个单词出现的次数，还有是什么单词
	var s = "how are you wq you"
	var wordCount = make(map[string]int, 10)
	words := strings.Split(s, " ")
	for _, word := range words {
		v, ok := wordCount[word]
		if ok {
			wordCount[word] = v + 1
		} else {
			wordCount[word] = 1
		}
	}
	for k, v := range wordCount {
		fmt.Println(k, v)
	}

	fmt.Println("练习题2")
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)

	fmt.Println("练习题3")
	type Map map[string][]int
	m := make(Map)
	s3 := []int{1, 2}
	s3 = append(s3, 3)
	fmt.Printf("%+v\n", s3)
	m["q1mi"] = s3
	s3 = append(s3[:1], s3[2:]...)
	fmt.Printf("%+v\n", s3)
	fmt.Printf("%+v\n", m["q1mi"])
}
