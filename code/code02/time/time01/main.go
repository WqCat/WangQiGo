package main

import (
	"fmt"
	"time"
)

//time demo
func main() {
	now := time.Now() //时间对象
	fmt.Println(now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Println(year, month, day, hour, minute, second)
	//时间戳：从1970.1.1到现在的秒数
	timeStamp1 := now.Unix()     //秒数
	timeStamo2 := now.UnixNano() //纳秒数
	fmt.Println(timeStamp1, timeStamo2)
	//将时间戳转换为具体的时间格式
	//1651543957+3600
	t := time.Unix(1651543957, 0)
	fmt.Println(t)

	//时间间隔
	// n := 5
	// time.Sleep(time.Duration(n) * time.Second)
	// fmt.Println("over")

	//time 的Add
	fmt.Println(now)
	t2 := now.Add(time.Hour)
	fmt.Println(t2)
	//Sub
	fmt.Println(t2.Sub(now))

	//定时器
	// for tmp := range time.Tick(time.Second * 3) {
	// 	fmt.Println(tmp)
	// }

	//时间格式化 Y   - m-  d   H : M: S
	//          2006  01  02  15  04  05
	ret1 := now.Format("2006/01/02") //2006.01.02也可
	fmt.Println(ret1 + "***************")
	ret2 := now.Format("2006-01-02 15:04:05")
	fmt.Println(ret2)
	ret3 := now.Format("2006-01-02 03:04:05.000 PM")
	fmt.Println(ret3)

	//解析字符串类型的时间
	timeStr := "2020/08/07 10:38:20"

	//根据时区去解析一个字符串格式时间
	timeObj, err := time.Parse("2006/01/02 15:04:05", timeStr)
	if err != nil {
		fmt.Printf("parse timeStr failed,err:%v\n", err)
		return
	}
	fmt.Println(timeObj)

	//1.拿到时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	//2.根据时区去解析一个字符串格式的时间
	timeObj2, err := time.ParseInLocation("2006/01/02 15:04:05", timeStr, loc)
	if err != nil {
		fmt.Printf("parse timeStr failed,err:%v\n", err)
		return
	}
	fmt.Println(timeObj2)

}
