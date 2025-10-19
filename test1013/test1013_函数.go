package main

import "fmt"

//求俩个数字的和

/*
func 函数名(参数)(返回值){
	函数体
}

*/
func sumFn(x int, y int) int {
	sum := x + y
	return sum
}

func sumFN(x, y int) int {
	return x + y
}

//求俩个数差（输入俩个都是相同类型参数，参数类型写一个即可）
func subFn(x, y int) int {
	sub := x - y
	return sub
}

//函数的可变参数
func sumFn1(x ...int) int {

	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum

}

func sumFn2(x int, y ...int) int { //x= 100,y  2, 2, 2

	sum := x
	for _, v := range y {
		sum += v
	}
	return sum

}

//一次返回多个值

func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

// 返回值命名

func calc1(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

func test() {
	fmt.Println("执行测试-----")
}

func main() {
	sum1 := sumFn(12, 123)
	fmt.Println(sum1)

	sub1 := subFn(22, 9)
	fmt.Println(sub1)

	sum2 := sumFn1(1, 2, 3, 4, 5, 6, 7, 8888)
	fmt.Println(sum2)

	sum3 := sumFn1(100, 2, 2, 2)
	fmt.Println(sum3)

	num1, num2 := calc(12, 33)
	fmt.Println(num1, num2)

	num3, num4 := calc1(44, 44)
	fmt.Println(num3, num4)

	//也可以只接受一个值    用_代替不接受的值
	_, num5 := calc1(44, 44)
	fmt.Println(num5)

	test()
}
