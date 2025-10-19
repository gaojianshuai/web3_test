package main

import (
	"fmt"
	"unsafe"
)

// 全局指针变量
var i int
var p1 *int = &i

var g string = "hello"
var p2 *string = &g

func main() {

	i := 1
	s := "hello"

	// 基础类型数据，必须使用变量名获取指针，无法通过字面量获取指针，字面量也就是直接就是字符串类似

	// 因为字面量会在编译期被声明成常量，不能获取内存中的指针信息
	//局部指针变量
	p1 = &i
	p2 = &s

	//var p3 **string = &p2  指针的指针 ：获取的是指针变量的内存地址
	p3 := &p2
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	// 指针也是有零值得，他的零值是nil，比如：var p4 *byte
	var p4 *byte
	//零值
	fmt.Println(p4)
	// 有内存地址输出
	fmt.Println(&p4)

	//指针于指针之间不能参与计算
	// fmt.Println(&p4 + &p3)

	//如何使用指针
	fmt.Println(*p1)
	fmt.Println(*p2)
	fmt.Println(**p3)
	// fmt.Println(*p4)

	//修改指针指向的值

	m := 2
	var n *int
	fmt.Println(&m)
	n = &m
	fmt.Println(n, &m)

	var mm **int
	mm = &n
	fmt.Println(mm, n)

	**mm = 5
	fmt.Println(mm, *mm, n)
	fmt.Println(**mm, *n)
	fmt.Println(m, &m)

	// 指针偏移操作unsafe.Point   和uintptr   不推荐  不可控
	a := "hello, world"
	upA := uintptr(unsafe.Pointer(&a))
	upA += 1

	c := (*uint8)(unsafe.Pointer(upA))
	fmt.Print(*c)

}
