package main

import "fmt"

func main() {
	// 1.创建管道：先入先出
	ch := make(chan int, 3)

	//2.给管道里面存储数据
	ch <- 10
	ch <- 20
	ch <- 33
	//3.获取管道里面的内容

	a := <-ch
	fmt.Println(a)

	<-ch      //从管道里取值  20
	c := <-ch //32
	fmt.Println(c)
	//4.打印管道的长度和容量
	ch <- 99
	fmt.Printf("值：%v  容量： %v  长度：%v\n", ch, cap(ch), len(ch)) //值：0xc00006e000  容量： 3  长度：1

	//5.管道的类型：引用数据类型
	ch1 := make(chan int, 4)
	ch1 <- 11
	ch1 <- 22
	ch1 <- 33

	ch2 := ch1
	ch2 <- 44
	<-ch1
	<-ch1
	<-ch1

	d := <-ch1
	fmt.Println(d)

	//管道阻塞：管道放满放不下fatal error: all goroutines are asleep - deadlock!
	// ch4 := make(chan int, 1)
	// ch4 <- 12
	// ch4 <- 54 //fatal error: all goroutines are asleep - deadlock!

	ch5 := make(chan string, 2)
	ch5 <- "java"
	ch5 <- "golang"

	// m1 := <-ch5
	// m2 := <-ch5
	// m3 := <-ch5
	// fmt.Println(m1, m2, m3) //fatal error: all goroutines are asleep - deadlock!

	//正确的写法就是：一遍写入 一遍出

	ch6 := make(chan string, 2)
	ch6 <- "java"
	<-ch6
	ch6 <- "golang"
	<-ch6

	//for range和for循环来遍历管道：range遍历必须加close()函数，for循环不需要加

	var ch7 = make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch7 <- i
	}
	close(ch7)

	//管道没有key，所以直接value就行
	for v := range ch7 {
		fmt.Println(v)
	}

	//for循环遍历:for循环管道的时候，管道可以不关闭
	var ch8 = make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch8 <- i
	}

	//管道没有key，所以直接value就行
	for j := 0; j < 10; j++ {
		fmt.Println(ch8)
	}

}
