package main

import "fmt"

//数组相关内容
func main() {
	var a [3]int
	var b [4]int
	fmt.Println(a)
	fmt.Println(b)
	//初始化
	//1.定义时使用初始化值
	var cityArray = [4]string{"徐州", "南京", "上海", "南通"}
	fmt.Println(cityArray)
	fmt.Println(cityArray[0])
	//2.编译器推导数组长度
	var boolArray = [...]bool{true, false, true, true, false}
	fmt.Println(boolArray)
	//3.使用索引方式初始化
	var langArray = [...]string{1: "Golang", 3: "Python", 7: "Java"}
	fmt.Println(langArray)
	fmt.Printf("%T\n", langArray)

	//数组遍历
	//1.for循环
	for i := 0; i < len(cityArray); i++ {
		fmt.Println(cityArray[i])
	}
	//2.for range遍历
	for _, value := range cityArray {
		fmt.Println(value)
	}

	//二维数组
	fmt.Println("二维数组")
	city2Array := [...][2]string{
		{"北京", "徐州"},
		{"南京", "南通"},
		{"镇江", "无锡"},
		{"济南", "连云港"},
	}
	fmt.Println(city2Array)
	fmt.Println(city2Array[0][1])
	//二维数组遍历
	fmt.Println("二维数组遍历")
	for _, v1 := range city2Array {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	//数组是值类型
	x := [...][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(x)
	f1(x)
	fmt.Println(x)
	y := x
	y[0][0] = 1000
	fmt.Println(x)

}
func f1(a [3][2]int) {
	a[0][0] = 100
}
