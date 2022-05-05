package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//文件操作

func readFromFile() {
	fileObj, err := os.Open("./wq.txt") //相对路径，同目录
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}
	defer fileObj.Close() //关闭文件
	//读取文件的内容
	for {
		var temp = make([]byte, 128)
		n, err := fileObj.Read(temp)
		if err == io.EOF { //End Of File
			//把当前读了多少字节打印出来，然后退出
			fmt.Println(string(temp[:n]))
			return
		}
		if err != nil {
			fmt.Printf("read from file failed,err:%v\n", err)
			return
		}
		fmt.Printf("read %d bytes from file\n", n)
		fmt.Println(string(temp[:n]))
	}
}

//read by bufio
func readByBufio() {
	//打开文件
	fileObj, err := os.Open("./wq.txt")
	if err != nil {
		fmt.Printf("open file failed,err%v\n", err)
		return
	}
	defer fileObj.Close()
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("All file have read")
			return
		}
		if err != nil {
			fmt.Printf("read file by bufio failed,err%v\n", err)
			return
		}
		fmt.Println(line)
	}
}

//read file ioutil
func readByIoutil() {
	content, err := ioutil.ReadFile("./wq.txt")
	if err != nil {
		fmt.Printf("read file by ioutil failed,err:%v\n", err)
		return
	}
	fmt.Println(string(content))
}

//write
func write() {
	fileObj, err := os.OpenFile("./wq2.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644) //APPEND:追加  TRUNC：清空   WRONLY:只写  EDONLY: 只读  RDWR：读写
	if err != nil {
		fmt.Printf("open file failed err:%v\n", err)
		return
	}
	defer fileObj.Close()
	str := "侠岚"
	fileObj.Write([]byte(str))      //[]byte
	fileObj.WriteString("hello 侠岚") //string
}

//write bufio
func writeByBufio() {
	fileObj, err := os.OpenFile("./wq.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file failed err:%v\n", err)
		return
	}
	defer fileObj.Close()
	writer := bufio.NewWriter(fileObj)
	writer.WriteString("山鬼谣") //将内容写入缓冲区
	writer.Flush()            //将缓冲区写入磁盘
}

//write ioutil
func writeByIoutil() {
	str := "伸出你的左手,看看手中有没有一个印记，如果有，也许世界的命运将掌握在你的手中。--侠岚"
	err := ioutil.WriteFile("./wq.txt", []byte(str), 0644)
	if err != nil {
		fmt.Printf("write file failed,err:%v\n", err)
		return
	}
}

func main() {
	// readFromFile()
	// readByBufio()
	// readByIoutil()
	// // write()
	// writeByBufio()
	writeByIoutil()
}
