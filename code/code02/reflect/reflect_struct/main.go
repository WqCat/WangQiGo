package main

import (
	"fmt"
	"reflect"
)

//结构体反射

type student struct {
	Name  string `json:"name" ini:"n_name"`
	Score int    `json:"score" ini:"s_score"`
}

// 给student添加两个方法 Study和Sleep(注意首字母大写)
func (s student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
	// 通过方法名获取结构体的方法
	// t.MethodByName("Sleep")  //(Method,bool)
	methodObj := v.MethodByName("Sleep") //value
	fmt.Println(methodObj.Type())
}

func main() {
	stu1 := student{
		Name:  "盖聂",
		Score: 90,
	}
	//通过反射区获取结构体中的所有字段信息
	t := reflect.TypeOf(stu1)
	fmt.Printf("name:%v kind:%v\n", t.Name(), t.Kind())
	//遍历结构体变量的所有字段
	for i := 0; i < t.NumField(); i++ {
		//根据结构体字段的索引去取字段
		fileObj := t.Field(i)
		fmt.Printf("name:%v type:%v tag:%v\n", fileObj.Name, fileObj.Type, fileObj.Tag)
		fmt.Println(fileObj.Tag.Get("json"), fileObj.Tag.Get("ini"))
	}

	//根据名字取结构体中的字段
	fmt.Println("根据名字取结构体中的字段")
	fileObj2, ok := t.FieldByName("Score")
	if ok {
		fmt.Printf("name:%v type:%v tag:%v\n", fileObj2.Name, fileObj2.Type, fileObj2.Tag)
	}

	printMethod(stu1)

}
