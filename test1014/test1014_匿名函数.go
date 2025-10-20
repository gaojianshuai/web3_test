package main

import "fmt"

func main() {
	//匿名函数
	func() {
		fmt.Println("test11")
	}()

	var fn = func(x, y int) int {
		return x + y
	}

	fmt.Println(fn(4, 4))

	//匿名函数自执行函数接受参数
	func(x, y int) {
		fmt.Println(x + y)
	}(33, 44)

}
