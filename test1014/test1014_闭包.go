package main

import "fmt"

/*
golang闭包

全局变量特点：
	1.常驻内存
	2.污染全局
局部变量特点：
	1.不常驻内存
	2.不污染全局

闭包：
	1.可以让一个变量常驻内存
	2.可以让一个变量不污染全局

闭包：
	1.闭包是只有权访问领一个函数作用域中的变量的函数
	2.创建闭包的常见方式就是在一个函数内部创建另一个函数，通过另一个函数访问这个函数的局部变量


*/

//写法：闭包的写法，函数里面嵌套一个函数，最后返回里面的函数
func adder() func() int {
	var i = 10
	return func() int {
		return i + 1
	}
}

func adder1() func(y int) int {
	var i = 10
	return func(y int) int {

		i += y
		return i
	}
}

func main() {

	var fn = adder()
	fmt.Println(fn())

	var fn1 = adder1()
	fmt.Println(fn1(10))
	fmt.Println(fn1(10))
	fmt.Println(fn1(10))
}
