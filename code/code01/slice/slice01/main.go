package main

import "fmt"

//切片（slice)
func main() {
	//基于数组得到切片
	a := [5]int{55, 56, 57, 58, 59}
	b := a[1:4]
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	//切片可以接着切片

	//make函数构造切片
	d := make([]int, 5, 10)
	fmt.Println(d)
	fmt.Printf("%T\n", d)
	//通过Len获取切片长度
	fmt.Println(len(d))
	//通过cap(）函数获取切片的容量
	fmt.Println(cap(d))

	//nil
	var a1 []int     //声明int类型切片
	var b1 = []int{} //声明并且初始化
	c := make([]int, 0)
	if a1 == nil {
		fmt.Println("a==nil")
	}
	fmt.Println(a1, len(a1), cap(a1))
	if b1 == nil {
		fmt.Println("b1==nil")
	}
	fmt.Println(b1, len(b1), cap(b1))
	if c == nil {
		fmt.Println("c==nil")
	}
	fmt.Println(c, len(c), cap(c))

	//切片b:=a 吧b[0]=100,a也会变

	//切片遍历
	e := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(e); i++ {
		fmt.Println(i, e[i])
	}
	fmt.Println()
	for index, value := range e {
		fmt.Println(index, value)
	}
	fmt.Println("扩容")
	//切片的扩容append
	var r []int
	var r1 []int
	r1 = append(r1, 10)
	fmt.Println(r1)
	// for i := 0; i < 10; i++ {
	// 	r = append(r, i)
	// 	fmt.Printf("%v len:%d cap:%d ptr:%p\n", r, len(r), cap(r), r)
	// }
	r = append(r, 1, 2, 3, 4, 5)
	fmt.Println(a)
	s := []int{12, 13, 45, 46}
	r = append(r, s...)
	fmt.Println(r)

}
