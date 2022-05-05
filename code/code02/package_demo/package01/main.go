package package01

import (
	"code/code02/package_demo/snow"
	"fmt"
)

//标识符首字母大写，表示可见
//通常不对外的标识符都是要首字母小写

//Name 定义一个测试的全局变量
var Name = "秦时明月"

//Add 是一个计算和的函数
func Add(x, y int) int {
	snow.Snow()
	Sub(x, y) //同一个包的直接用
	return x + y
}

//init函数在包导入的时候自动执行
//init函数没有参数也没有返回值
//在全局声明后
func init() {
	fmt.Println("package01.init()函数")
	fmt.Println(Name) //全局声明后
}
