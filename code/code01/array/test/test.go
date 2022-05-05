package main

import "fmt"

func main() {
	a := [5]int{1, 3, 5, 7, 8}
	b := [2][2]int{
		{0, 0},
	}
	flag := 0
	aLen := len(a)
	for i := 0; i < aLen-1; i++ {
		if a[i] > 8 {
			break
		}
		for j := aLen - 1; j > i; j-- {
			if a[j]+a[i] < 8 {
				break
			}
			if a[i]+a[j] == 8 {
				b[flag][0] = a[i]
				b[flag][1] = a[j]
				flag++
			}
		}
	}
	// for _, v1 := range b {
	// 	for _, v2 := range v1 {
	// 		fmt.Println(v2)
	// 	}
	// }
	fmt.Println(b)
}
