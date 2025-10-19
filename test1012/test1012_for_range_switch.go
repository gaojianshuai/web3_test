package main

import (
	"fmt"
)

func main() {

	// go语言中，for range遍历数组、切片、字符串、map、及通道channel

	var str = "你好golang"

	for k, v := range str {
		fmt.Printf("key=%v val=%v \n", k, v)

	}

	var arr = []string{"golang", "java", "python", "node", "javascript"}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for _, val := range arr {
		fmt.Println(val)
	}

	// switch case:练习：判断文件类型后缀是.html 输出text/html   如果后缀是.css 输出text/css  如果是.js
	//输出text/javascript

	var ext = ".html"
	if ext == ".html" {
		fmt.Println("text/html")
	} else if ext == ".css" {
		fmt.Println("text/css")
	} else if ext == ".js" {
		fmt.Println("text/javascript")
	} else {
		fmt.Println("找不到此格式文件")
	}

	//1.switch case得基本用法：break可以写也可以不写，建议写这
	var extname = ".css"

	switch extname {
	case ".html":
		fmt.Println("text/html")
		break
	case ".css":
		fmt.Println("text/css")
		break
	case ".js":
		fmt.Println("text/javascript")
		break
	default:
		fmt.Println("找不到此格式文件")
	}

	//2.switch case得第二种用法

	switch extname1 := ".css"; extname1 {
	case ".html":
		fmt.Println("text/html")
		break
	case ".css":
		fmt.Println("text/css")
		break
	case ".js":
		fmt.Println("text/javascript")
		break
	default:
		fmt.Println("找不到此格式文件")
	}

	//3. 一个分支可以有多个值，多个case值中间使用英文逗号分隔

	var mm = 9
	switch mm {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
		break
	case 2, 4, 6, 8, 10:
		fmt.Println("偶数")
		break

	}

	// abc及格。d 不及格
	score1 := "A"
	switch score1 {
	case "A", "B", "C":
		fmt.Println("及格")
		break
	case "D":
		fmt.Println("不及格")
		break
	}

	//4.分支还可以使用表达式，这时候，siwtch 语句后面不需要跟判断变量
	var age = 30
	switch {
	case age < 24:
		fmt.Println("好好学习")
	case age >= 24 && age <= 60:
		fmt.Println("好好赚钱")
	case age > 60:
		fmt.Println("注意身体")
	default:
		fmt.Println("输入错误")
	}

	//5.switch 得穿透  fallthrought:即输出好好赚钱，也输出注意身体
	var age1 = 30
	switch {
	case age1 < 24:
		fmt.Println("好好学习")
	case age1 >= 24 && age1 <= 60:
		fmt.Println("好好赚钱")
		fallthrough
	case age1 > 60:
		fmt.Println("注意身体")
	default:
		fmt.Println("输入错误")
	}
}
