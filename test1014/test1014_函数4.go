package main

import "fmt"

type calc func(int, int) int //表示顶一个一个calc的类型

type myInt int

func add1(x, y int) int {
	return x + y
}

func sub1(x, y int) int {
	return x - y
}

// 函数作为另一个函数的参数
// 自定义一个方法类型
type calcType func(int, int) int

func calc1(x, y int, cb calcType) int {
	return cb(x, y)
}

//函数作为返回值

type calcType2 func(int, int) int

func do(o string) calcType2 {
	switch o {
	case "+":
		return add1
	case "-":
		return sub1
	case "*":
		return func(x, y int) int {
			return x * y
		}
	default:
		return nil
	}
}

func main() {
	s := calc1(4, 5, add1)
	fmt.Println(s)

	j := calc1(3, 4, func(x, y int) int {
		return x * y
	})
	fmt.Println(j)

	var d = do("+")
	fmt.Println(d(6, 7))

	e := do("*")
	fmt.Println(e(5, 5))
}
