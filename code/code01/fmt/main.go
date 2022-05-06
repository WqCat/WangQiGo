package main

import "fmt"

//fmt
func main() {
	fmt.Println("徐州")
	//Printf("格式化字符串",值)
	//%T:查看类型
	//%b %o %d %x十六进制(a-f)   %X十六进制(A-F)
	//%c字符
	//%s:字符串
	//%p:指针
	//%v:值
	//%f:浮点数
	//%t 布尔型
	var m = make(map[string]int, 1)
	m["徐州"] = 100
	fmt.Printf("%v\n", m)
	fmt.Printf("%#v\n", m)

	printBaifenbi(90)
	fmt.Printf("%q\n", 65) //ASCII码
}
func printBaifenbi(num int) {
	fmt.Printf("%d%%\n", num) //%%-->%
}
