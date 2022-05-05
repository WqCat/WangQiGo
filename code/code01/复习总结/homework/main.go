package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	//1.判断字符串中汉字的数量
	//难点是判断是否是汉字
	fmt.Println("1.判断字符串中汉字的数量")
	s1 := "Hello徐州"
	//1.依次拿到字符串中的字符
	var count int
	for _, v := range s1 {
		//2.判断是否是汉字
		if unicode.Is(unicode.Han, v) {
			count++
		}
	}
	//3.把汉字出现次数累加
	fmt.Println(count)
	//2.how do you do wq
	fmt.Println("2.how do you do wq")
	s2 := "how do you do wq"
	//2.1按照字符串切割
	s3 := strings.Split(s2, " ")
	//2.2遍历切片存储到一个map
	m1 := make(map[string]int, 10)
	for _, w := range s3 {
		//1.如果原来map中不存在v这个key，那么出现次数=1
		if _, ok := m1[w]; !ok {
			m1[w] = 1
		} else {
			m1[w]++
		}
		//2.如果map中存在v这个key，那么出现次数+1
	}
	//2.3累加出现的次数
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	//3.回文序列
	fmt.Println("3.回文序列")
	str := "十八到日本日到八十"
	// for i := 0; i < strLen; i++ {
	// 	fmt.Println(str[i])
	// }
	// for i := 0; i < strLen/2; i++ {
	// 	if str[i] != str[strLen-i-1] {         //str[i]是8进制码表示汉字
	// 		fmt.Println("不是回文序列")
	// 		return
	// 	}
	// }
	// fmt.Println("是回文")
	r := make([]rune, 0, len(str))
	for _, c := range str {
		r = append(r, c)
	}
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")
}
