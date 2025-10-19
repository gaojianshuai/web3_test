package main

import "fmt"

func fn4() {
	fmt.Println("测试开始111")
	defer func() {
		fmt.Println("延迟批量执行函数1")
		fmt.Println("延迟批量执行函数2")
		fmt.Println("延迟批量执行函数3")
	}()
}

//匿名返回值
func fn5() int {
	x := 5
	defer func() {
		x++
	}()
	return 5
}

//命名返回值

func fn6() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func main() {
	//1.defer得基本使用

	fmt.Println("测试执行")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("测试结束")

	//2.defer在命名返回值和匿名返回 函数中表现不一样
	fn4()
	//匿名返回值
	fmt.Println(fn5())

	//命名返回值
	fmt.Println(fn6())

}
