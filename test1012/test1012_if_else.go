package main

import "fmt"

func main() {

	/*
		if 表达式1{
			分支1
		}else if 表达式2 {
			分支2
		}else {
			分支3
		}

	*/

	flag := true

	if flag {
		fmt.Println(flag)
	}
	//第一种if写法：相当于当前区域的全局变量age
	age := 35
	if age > 30 {
		fmt.Println("三十而立")
	}
	fmt.Println(age)

	//if语句的另一种写法：这里的age相当于局部变量

	if age1 := 34; age1 > 30 {
		fmt.Println("成年人")
	}
	// fmt.Println(age1)
	// 输入一个人的成绩。如果成绩大于等于90 输出A    小于90大于75输出B  否则输出c
	score := 77
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	//另一种写法：

	if score := 60; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	// 1.if else得细节：if{}不能省略
	// if age1 := 34; age1 > 30 {
	// 	fmt.Println("成年人")
	// }

	//2.{}必须紧挨着条件，括号不能换行

	//求俩个数的最大值

	var a = 23
	var b = 66
	var max int

	if a > b {
		max = a
	} else {
		max = b
	}
	fmt.Println(max)

	/*
		for 初始语句;条件表达式:结束语句{
			循环体语句
		}

	*/

	// 打印1-10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	//第二种写法：注意for后面的冒号
	j := 1
	for ; j <= 10; j++ {
		fmt.Println(j)
	}
	//第三种写法
	x := 1
	for x <= 10 {
		fmt.Println(x)
		x++
	}

	//第四种：无限循环
	y := 1
	for {
		if y <= 10 {
			fmt.Println(y)
		} else {
			break
		}
		y++
	}

	// golang中没有while语句，用for代替、

	// 练习1：打印0-50的偶数

	for z := 0; z <= 50; z++ {
		if z%2 == 0 {
			fmt.Println(z)
		}
	}
	//练习2：求1--100的和

	sum := 0
	for i := 1; i <= 100; i++ {
		sum = sum + i
	}
	fmt.Println(sum)

	//3. 打印1--100之间是9的倍数的总和

	sum1 := 0
	count := 0
	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			count = count + 1
			sum1 = sum1 + i
		}
	}
	fmt.Println(sum1)
	fmt.Println(count)

	// 练习4：计算5的阶乘    1*2*3*4*5

	var sum2 = 1
	for i := 1; i <= 5; i++ {
		sum2 = sum2 * i
	}
	fmt.Println(sum2)

	//5.打印一个*号矩形
	/*
	*******
	*******
	*******
	 */
	for i := 1; i <= 21; i++ {
		fmt.Print("*")
		if i%7 == 0 {
			fmt.Println("")
		}
	}

	//for循环嵌套

	var row = 5
	var column = 8
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

	//7. 打印一个三角形

	var row1 = 5
	for i := 1; i <= row1; i++ {

		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

	//8.打印 99乘法口诀表

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v*%v=%v \t", i, j, i*j)
		}
		fmt.Println("")
	}
}
