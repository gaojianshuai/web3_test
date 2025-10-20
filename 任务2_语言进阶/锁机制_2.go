package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func main() {
	var counter int64 // 使用int64类型的原子计数器
	var wg sync.WaitGroup

	// 启动10个协程
	numGoroutines := 10
	iterations := 1000

	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()

			// 每个协程执行1000次递增操作
			for j := 0; j < iterations; j++ {
				atomic.AddInt64(&counter, 1) // 原子递增操作
			}
		}()
	}

	// 等待所有协程完成
	wg.Wait()

	// 安全地读取最终值
	finalValue := atomic.LoadInt64(&counter) // 原子加载操作
	fmt.Printf("最终计数器值: %d\n", finalValue)

	// 验证结果
	expected := int64(numGoroutines * iterations)
	fmt.Printf("期望值: %d\n", expected)

	if finalValue == expected {
		fmt.Println("✓ 计数器操作正确")
	} else {
		fmt.Println("✗ 计数器操作存在错误")
	}
}
