package main

import (
	"fmt"
	"time"
)

func main() {
	timeObj := time.Now()
	fmt.Println(timeObj) //2025-10-14 13:53:10.08788 +0800 CST m=+0.000523201

	//要求输出2025-10-14 13:53:10
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()

	fmt.Printf("%d-%d-%d %d:%d:%d\n", year, month, day, hour, minute, second) //2025-10-14 13:55:57
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	//format方法格式化时间
	timeObj1 := time.Now()
	formattedTime := timeObj1.Format("2006-01-02 03:04:05")
	fmt.Println(formattedTime)

	timeObj2 := time.Now()
	formattedTime1 := timeObj2.Format("2006/01/02 15:04:05")
	fmt.Println(formattedTime1)

	//获取当前时间戳
	timeObj3 := time.Now()
	unixtime := timeObj3.Unix() //获取当前时间戳：1970年到现在的毫秒时间
	fmt.Println(unixtime)

	//获取当前纳秒时间戳
	unixNaTime := timeObj3.UnixNano()
	fmt.Println(unixNaTime)

	// 时间戳转换成日期字符串  1760422212
	unitTime1 := 1760422212
	timeObj = time.Unix(int64(unitTime1), 0)
	str1 := timeObj.Format("2006-01-02 03:04:05")
	fmt.Println(str1)

	//字符串转换成时间戳
	var tmp = "2025-05-01 00:00:00"
	var str2 = "2025-10-14 02:10:12"

	timeObj4, _ := time.ParseInLocation(tmp, str2, time.Local)
	fmt.Println(timeObj4)

	fmt.Println(time.Millisecond)
	fmt.Println(second)

	//时间操作函数
}
