package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 为了保证协程能执行完毕后退出，这个时候我们可以使用sync.WaitGroup等待协程执行完毕
var wg sync.WaitGroup

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test 你好,golang")
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() //调用完成后 协程计数器-1
}

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println("test1 你好,golang")
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() //协程计数器-1
}

func main() {

	//进程：一个正在执行的程序   线程：是进程的一个执行实例
	// 并发：多个线程同时竞争一个位置，竞争打的才可以执行，每个时间段只有一个线程在执行
	// 并行：多个线程可以同时执行，每一个时间段可以有多个线程同时执行

	// 主线程可以理解成进程中，开启一个goroutine，该携程每哥50毫秒 输出 你好golang
	// 在主线程中每隔50毫秒输出你好golang 输出10次后， 退出程序
	//注意：当主进程执行结束后，不管协程结束没结束，都会直接退出 可以使用sync.waitGroup 等待协程结束后退出
	wg.Add(1)  //协程计数器+1
	go test()  //表示开启一个协程
	wg.Add(1)  //协程计数器+1
	go test1() //表示开启一个协程
	for i := 0; i < 10; i++ {
		fmt.Println("你好,golang")
		time.Sleep(time.Millisecond * 100)
	}

	//time.Sleep(time.Second)
	wg.Wait() //等待协程执行完毕后
	fmt.Println("主线程退出")

	// 获取当前几岁按期上面的cpu个数
	cpuNum := runtime.NumCPU()
	fmt.Println(cpuNum)

	// 可以自己设置使用多个cpu
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")
}
