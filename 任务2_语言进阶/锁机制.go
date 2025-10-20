package main

import (
	"fmt"
	"sync"
)

// 难度：中等，不太熟这块
// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

	// 等待所有协程完成
	wg.Wait()

	// 输出最终结果
	fmt.Printf("最终计数器值: %d\n", counter.Value())
	fmt.Printf("期望值: 10000\n")
	if counter.Value() == 10000 {
		fmt.Println("✓ 结果正确，并发安全得到保障")
	} else {
		fmt.Println("✗ 结果错误，存在并发问题")
	}
}
