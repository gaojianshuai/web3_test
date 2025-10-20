package main

import (
	"fmt"
	"sync"
)

// 实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
func main() {
	// 创建带缓冲的通道，缓冲大小为10
	ch := make(chan int, 10)

	// 使用WaitGroup等待所有goroutine完成
	var wg sync.WaitGroup

	// 生产者协程
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch) // 生产完成后关闭通道

		for i := 1; i <= 100; i++ {
			ch <- i
			fmt.Printf("生产者发送: %d\n", i)
		}
		fmt.Println("生产者完成所有数据发送")
	}()

	// 消费者协程
	wg.Add(1)
	go func() {
		defer wg.Done()

		for num := range ch {
			fmt.Printf("消费者接收: %d\n", num)
		}
		fmt.Println("消费者完成所有数据接收")
	}()

	// 等待所有协程完成
	wg.Wait()
	fmt.Println("程序执行完毕")
}
