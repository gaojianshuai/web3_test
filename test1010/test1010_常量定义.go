package main


import "fmt"


func main(){

	//申明常量:常量申明之后不能改变
	const gaojs = "gaojianshuai"
	fmt.Println(gaojs)

	//声明多个常量
	const (
		a = 111
		b = "ttt"
	)
	fmt.Println(a, b)
	// const同时申明多个常量，且后面的值省略的话，则和第一个值一样
	const (
		c = 111
		d
	)
	fmt.Println(c, d)

	//const和iota得使用
	const a2 = iota
	fmt.Println(a2)
	// iota初始化常量为0，入伙后面常量缺省赋值，则从零递增
	const (
		f = iota	//0
		g	//1
		h	//2
	)
	fmt.Println(f, g, h)
	//可以跳过匿名常量值
	const (
		l = iota	//0
		_	//1
		n	//2
	)
	fmt.Println(l, n)

	//iota可以插队

	const (
		n1 = iota	//0
		n2 = 100	//
		n3= iota//
		n4
	)
	fmt.Println(n1, n2, n3, n4)

	//多个iota可以定义一行

}