package main

import "fmt"
import "unsafe"
// import "github.com/shopspring/decimal"

func main(){

	//1.浮点型
	var num float64 = 4.4
	fmt.Printf("num=%v 数据类型是%T\n", num, num)
	fmt.Println(unsafe.Sizeof(num))
	//默认%f保留6 位小数
	var num1 float64 = 3.1415926535
	fmt.Printf("%v---%f\n", num1, num1)

	//默认%.2f保留2位小数
	fmt.Printf("%v---%.2f\n", num1, num1)

	//go语言科学计数法表示浮点型biaoshi 3.14乘以10的二次方
	var m1 float64 = 3.14e2
	fmt.Printf("%v---%T\n", m1, m1)

	//go语言科学计数法表示浮点型biaoshi 3.14除以10的二次方
	var m2 float64 = 3.14e-2
	fmt.Printf("%v---%T\n", m2, m2)

	//float精度丢失问题：二进制精度转换典型精度丢失问题112959.99999999999
	var m3 float64 = 1129.6
	fmt.Println(m3 * 100)
	//解决方法：使用第三方包：decimal包


	//整形和浮点型转换
	a := 10
	b := float64(a)
	fmt.Printf("a=%v得类型%T, b=%v得类型%T\n", a,a,b, b)

	//go语言中：bool型无法参加运算，也不能强制转换成bool型
	var flag = true
	fmt.Printf("%v----%T\n", flag, flag)

	//string变量的默认值为空
	var t1 string
	fmt.Printf("%v---%T\n",t1, t1)
	//int变量的默认值为0
	var t int
	fmt.Printf("%v---%T\n",t, t)

}