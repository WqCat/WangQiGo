package main

import "fmt"

//new-make
func main() {
	//new
	a := new(int)
	fmt.Println(*a)
	*a = 100
	fmt.Println(*a)
	//make
	b := make(map[string]int, 10)
	b["wq"] = 100
	fmt.Println(b)

}
