package main

import (
	"fmt"
	"io"
	"os"
)

//CopyFile 拷贝文件函数
func CopyFile(dsName, srcName string) (written int64, err error) {
	//读的方式打开文件
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("open %s failed err:%v\n", srcName, err)
		return
	}
	defer src.Close()
	//以写|创建的方式打开目标文件
	fileObj, err := os.OpenFile(dsName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open %s file failed,err:%v\n", dsName, err)
		return
	}
	defer fileObj.Close()
	return io.Copy(fileObj, src) //调用io.Copy()拷贝内容
}
func main() {
	_, err := CopyFile("wqCopy.txt", "wq.txt")
	if err != nil {
		fmt.Printf("copy file failed,err:%v", err)
		return
	}
	fmt.Println("copy done!")
}
