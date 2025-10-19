package main
import "fmt"


var t = "全局变量，正常声明方法"

// 不能放在方法外：t1 := "全局变量：短变量申明法"



//6.匿名变量：在使用多重赋值时，如果想要忽略某个值，可以使用匿名变量

func getUserinfo()(string, int){
	return "李白", 99
}


func main(){

/*
	1.var生命变量
	2.var变量名称
	3.变量名称命名：G语言变量名由字母、数字、下划线组成，其中首个字母不能为数字
	go语言中关键字和保留字都不能作为变量名
	4.go语言中变量名字严格区分大小写
	5.变量名称采用驼峰命名规则
*/

	var username string
	fmt.Println(username)	//变量申明后灭有初始化的话--值为空

	var a1 = "张三"
	fmt.Println(a1)

	// 变量首个字母不能为数字
	// var 1a string
	// fmt.Printf(1a)

	var a_= "666"	//可以，但是不规范，不推荐
	fmt.Printf(a_)

	//go语言变量的定义和初始化:第一种
	var username1 string
	username1 = "gaojs"
	fmt.Println(username1)

	//第二种
	var username2 string = "777"
	fmt.Println(username2)

	//第三种
	var username3 = " 888"
	fmt.Println(username3)
	//同意作用域不能重复声明

	var username4 = "桃子"
	var age = 22
	var sex = "男"

	username4 = "苹果"	//二次赋值
	fmt.Println(username4, age, sex)

	// 2.一次申明多个变量:var 变量名称， 变量名称 类型相同

	var b1, b2, b3 string
	b1 = "香蕉"
	b2 = "葡萄"
	b3 = "猕猴桃"
	fmt.Println(b1, b2, b3)

	// 3. 一次申明多个不同类型的变量：首先申明在赋值，也可以直接赋值
	var (
		b4 string	// b4 string = "天安门"  或者b4 = "天安门"
		b5 int
		b6 string
		b7 bool
	)
	b4 = "水立方"
	b5 = 000
	b6 = "鸟巢"
	b7 = true
	fmt.Println(b4,b5,b6,b7)

	//4.短变量申明法：只能在当前作用域使用,只能使用短变量申明局部变量，不能申明全局变量

	b8 := "天坛"
	fmt.Println(b8)
	fmt.Printf("类型是：%T\n", b8)

	// 5.短变量申明多个变量，并且初始化变量
	x, y, z := 11, true, "阅兵"
	fmt.Println(x,y,z)
	fmt.Printf("x类型是：%T y类型是：%T z类型是：%T", x, y, z)

	// 打印全局变量
	fmt.Println(t)

	t1 := "全局变量：短变量申明法"
	fmt.Println(t1)

	//6.调用匿名函数:返回俩个值，如果我只要期中一个呢
	var username7, age1 = getUserinfo()
	fmt.Println(username7, age1)

	// 忽略一个值就可以这样写:使用匿名变量_,且不存在重复
	var username8, _ = getUserinfo()
	fmt.Println(username8)

	var _, age2 = getUserinfo()
	fmt.Println(age2)
}


// //6.匿名变量：在使用多重赋值时，如果想要忽略某个值，可以使用匿名变量

// func getUserinfo()(string, int){
// 	return "李白", 99
// }
