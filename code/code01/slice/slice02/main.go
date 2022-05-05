package main

import (
	"fmt"
	"sort"
)

func main() {
	//切片的copy
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5, 10)
	c := b
	copy(b, a)
	b[0] = 100
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//切片的删除
	a1 := []string{"徐州", "南京", "上海", "广州"}
	a1 = append(a1[0:2], a1[3:]...) //append(a[:index],[index+1:]...)
	fmt.Println(a1)
	var ssr = [...]int{3, 9, 5, 4, 7, 3, 2, 1} //数组
	//ssr[:]得到的是一个切片，指向了底层的数组a
	sort.Ints(ssr[:]) //这里面要是切片
	fmt.Println(ssr)
}
