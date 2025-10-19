package main

import (
	"fmt"
)

type calc2 func(int, int) int //表示顶一个一个calc的类型

type myInt int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func test() {
	fmt.Println("test111")
}
func main() {

	var c calc2
	c = add
	fmt.Printf("c得类型：%T", c)
	fmt.Println(c(2, 4))

	f := sub
	fmt.Printf("f得类型：%T\n", f)
	fmt.Println(f(5, 5))

	a := 10
	var b myInt = 54
	fmt.Printf("b得类型：%T\n", b)
	fmt.Println(a + int(b))
}
