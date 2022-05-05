package main

import (
	"fmt"
	"strings"
)

//字符串常见操作
func main() {
	//求字符串长度
	s := "hello"
	fmt.Println(len(s))
	//拼接字符串
	s1 := "wq"
	fmt.Println(s + s1)

	//字符串分割
	s2 := "how are you wq "
	fmt.Println(strings.Split(s2, " "))
	fmt.Printf("%T\n", strings.Split(s2, " "))
	//判断是否包含
	fmt.Println(strings.Contains(s2, "wq"))
	//判断前缀
	fmt.Println(strings.HasPrefix(s2, "how"))
	//判断后缀
	fmt.Println(strings.HasSuffix(s2, " "))
	//判断子串的位置
	fmt.Println(strings.Index(s2, "you"))
	//最后出现的位置
	fmt.Println(strings.LastIndex(s2, " "))
	//join
	s5 := []string{"how", "are", "you", "wq"}
	fmt.Println(s5)
	fmt.Println(strings.Join(s5, "+"))

}
