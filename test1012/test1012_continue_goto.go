package main

import "fmt"

func main() {

	//1. 表示i=2的时候跳出循环
	for i := 0; i < 10; i++ {
		if i == 2 {
			break
		}
		fmt.Println(i)
	}

	//2.档i=2的时候跳出循环：嵌套
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {

			if j == 3 {
				break
			}
			fmt.Println(i, j)
		}
	}

	//3.break在switch中执行一条case后跳出语句的作用

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

	//4.一次跳出俩个循环

label:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {

			if j == 3 {
				break label
			}
			fmt.Println(i, j)
		}
	}

	//continue语句可以结束当前循环，开始下一次的循环迭代过程，仅限在for循环内使用
	for i := 1; i <= 10; i++ {
		if i == 3 {
			continue //跳过i=3
		}
		fmt.Println(i)
	}

	//在continue语句后添加标签时，表示开始标签对应的循环

label2:
	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				continue label2
			}
			fmt.Println(i, j)
		}
	}

	//goto跳转到指定标签：goto语句通过标签尽心代码之间的无条件跳转
	var tt = 444
	if tt > 30 {
		fmt.Println("成年人")
		goto lebel3
	}

	fmt.Println("未成年")
	fmt.Println("中年人")
lebel3: //直接跳过了未成年和中年人
	fmt.Println("老年人")
	fmt.Println("小孩子")
	fmt.Println("大孩子")
}
