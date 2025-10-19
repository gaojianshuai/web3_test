package main

import (
	"fmt"
	"sync"
	"time"
)

// for循环开启多个goroutine ,开启10个携程，每个携程打印1-10
var wg1 sync.WaitGroup

func test11(num int) {
	defer wg1.Done()
	for i := 1; i <= 5; i++ {
		fmt.Printf("携程%v   打印的第%v条数据\n", num, i)
		time.Sleep(time.Millisecond * 100)
	}

}

func main() {
	for i := 1; i <= 6; i++ {
		wg1.Add(1)
		go test11(i)
	}
	wg1.Wait()
	fmt.Println("关闭主线程。。。")
}
