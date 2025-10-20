package main

import "fmt"

func main() {

	//1.有俩个变量，a  b   要求将其进行交换、
	var a = 10
	var b = 20
	//第一种
	var t = a
	a = b
	b = t

	fmt.Println(a, b)
	//第二种
	a = a + b //a = 10 +20 = 30
	b = a - b //b = 10+20 - 20 = 10
	a = a - b //a = 30 - 10=20
	fmt.Println(a, b)

	//练习  加入还有100天放假，问有几个星期零几天

	var days = 100
	var week = days / 7
	var day = days % 7
	fmt.Println(week, day)

	//练习：定义一个变量保存华氏温度，华氏温度转换摄氏温度的公式：
	//c = (F-32)/1.8           C代表摄氏温度   F代表华氏温度

	var F float32 = 100
	c := (F - 32) / 1.8
	fmt.Println(c)
	fmt.Printf("%.2f", c)

	//位运算符：对整数在二进制位操作

	var d = 5 //101
	var e = 2 //010

	fmt.Println("d&e=", d&e) //000
	fmt.Println("d|e=", d|e) //111
	fmt.Println("d^e=", d^e) //111 俩位不一样则为1
	//左移N位：就是乘以2得N次方  右移N位：就是除以2的N次方
	fmt.Println("d<<e=", d<<e) //5乘以2得2次方    20
	fmt.Println("d>>e=", d>>e) //5除以2的2次方    1

}
