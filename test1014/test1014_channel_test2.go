package main

import (
	"fmt"
	"sync"
	"time"
)

var wg4 sync.WaitGroup

func fb1(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("写入数据%v 成功\n", i)
		time.Sleep(time.Millisecond * 50)
	}
	close(ch)
	wg4.Done()
}

func fb2(ch chan int) {
	for v := range ch {
		fmt.Printf("读取数据%v 成功\n", v)
		time.Sleep(time.Millisecond * 50)
	}
	wg4.Done()
}

func main() {
	var ch = make(chan int, 10)

	wg4.Add(1)
	go fb1(ch)
	wg4.Add(1)
	go fb2(ch)

	wg4.Wait()

	fmt.Println("执行结束")
}

// 如果写入特别慢   读取特别慢，也不会影响也不会阻塞，会等待你写入完成。
