package main

import (
	"fmt"
	"sync"
)

// 为了保证协程能执行完毕后退出，这个时候我们可以使用sync.WaitGroup等待协程执行完毕
var wg sync.WaitGroup

func test() {
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println("test 协程打印奇数:", i)
		}

	}
	wg.Done() //调用完成后 协程计数器-1
}

func test1() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("test 协程打印偶数:", i)
		}

	}
	wg.Done() //调用完成后 协程计数器-1
}

func main() {
	wg.Add(1)  //协程计数器+1
	go test()  //表示开启一个协程
	wg.Add(1)  //协程计数器+1
	go test1() //表示开启一个协程
	wg.Wait()  //等待协程执行完毕后
	fmt.Println("主线程退出")

}
