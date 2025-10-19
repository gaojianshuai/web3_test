package main

import "fmt"

func fn1(k int) {
	if k > 0 {
		fmt.Println(k)
		k--
		fn1(k)
	}
}

func fn2(k int) int {
	if k > 1 {
		return k + fn2(k-1)
	} else {
		return 1
	}

}

func fn3(k int) int {
	if k > 1 {
		return k * fn3(k-1)
	} else {
		return 1
	}
}

//函数递归调用
func main() {
	//1.for循环实现1--100的和
	var sum = 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	//2.传入一个整数，递归打印出1-到这个数内所有整数
	fn1(5)

	//3.递归实现1--100的和
	fmt.Println(fn2(100))

	//4.函数的递归实现5的阶乘
	fmt.Println(fn3(5))
}
