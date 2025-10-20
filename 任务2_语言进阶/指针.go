package main

import (
	"fmt"
)

func numberPtrDemo(ptr *int) {
	// numberPtrDemo 接收一个整数指针，将其指向的值增加10
	if ptr != nil {
		*ptr = *ptr + 10
		fmt.Printf("函数内部：指针指向的值已增加10，当前值为：%d\n", *ptr)

	} else {
		fmt.Println("传入的数字指针为nil")

	}
}

func numberPtrDemo1(ptr *int) {
	// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
	if ptr != nil {
		*ptr = *ptr * 2
		fmt.Printf("函数内部：指针指向的值已增加10，当前值为：%d\n", *ptr)

	} else {
		fmt.Println("传入的数字指针为nil")

	}
}

func main() {
	// 初始化一个整数变量
	num := 25
	fmt.Printf("初始值：%d\n", num)

	// 调用函数，传入变量的地址
	numberPtrDemo(&num)
	fmt.Printf("主函数中：修改后的值为：%d\n", num)

	num1 := 50
	numberPtrDemo1(&num1)
	fmt.Printf("主函数中：修改后的值为：%d\n", num1)
}
