package main

import "fmt"

//panic和recover
func a() {
	fmt.Println("func a")
}

func b() {
	defer func() {
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来,可以让下面的程序继续运行
		if err != nil {
			fmt.Println("func b error,释放数据库连接。。。")
		}
	}()
	panic("panic in b") //程序崩溃，退出
}

func c() {
	fmt.Println("func c")
}
func main() {
	a()
	b()
	c()
}
