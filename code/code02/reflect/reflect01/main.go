package main

import (
	"fmt"
	"reflect"
)

//reflect

func reflectType(x interface{}) {
	//不知道别人调用时，会传进来什么类型的变量
	//1.类型断言
	//2.借助反射
	obj := reflect.TypeOf(x)
	fmt.Println(obj, obj.Name(), obj.Kind())
	fmt.Printf("%T\n", obj) //*reflect.rtype
}

type Cat struct{}

type Dog struct{}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("%v %T\n", v, v)
	k := v.Kind() //拿到值对应的类型种类
	//如何得到一个传入时候类型的变量
	switch k {
	case reflect.Float32:
		//将反射取到的值转换成一个int32类型的变量
		ret := float32(v.Float())
		fmt.Println(ret)
		fmt.Printf("%v %T\n", ret, ret)
	case reflect.Int32:
		ret := int32(v.Int())
		fmt.Println(ret)
		fmt.Printf("%v %T\n", ret, ret)
	}
}

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	//Elem()用来根据指针去对应的值
	k := v.Elem().Kind()
	switch k {
	case reflect.Int32:
		v.Elem().SetInt(100)
	case reflect.Float32:
		v.Elem().SetFloat(3.1433223)
	}
}

func main() {
	var a float32 = 1.234
	reflectType(a)
	var b int8 = 10
	reflectType(b)
	//结构体类型
	fmt.Println("结构体")
	var c Cat
	reflectType(c)
	//切片
	fmt.Println("切片")
	var e []int
	reflectType(e)
	var f []string
	reflectType(f)

	//valueOf
	fmt.Println("reflect取值")
	var aa int32 = 300
	reflectValue(aa)
	var bb float32 = 3.1456
	reflectValue(bb)
	//set Value
	fmt.Println("reflect设置值")
	var aaa int32 = 213
	reflectSetValue(&aaa)
	fmt.Println(aaa)
}
