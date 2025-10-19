package main

import "fmt"

//只接收channel的函数

func receiveOnly(ch <-chan int) {
	for v := range ch {
		fmt.Println("接收到：%d\n", v)
	}
}

// 只发送channel的函数

func sendOnly(ch chan<- int) {
	for i := 0; i < 6; i++ {
		ch <- i
		fmt.Println("发送：%d\n", i)
	}
	close(ch)
}
