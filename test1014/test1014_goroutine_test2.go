package main

import (
	"fmt"
	"sync"
	"time"
)

//要求：要统计1-120000的数字钟哪些是素数:分为四个协程1---30000  30001--60000  60001--90000  90001--120000
// start:(n-1)*30000+1  end:n*30000

var wg3 sync.WaitGroup

func test2(n int) {
	for num := (n-1)*3000 + 1; num <= n*30000; num++ {
		if num > 1 {
			var flag = true
			for i := 2; i < num; i++ {
				if num%i == 0 {
					flag = false
					break
				}
			}
			if flag {
				// fmt.Println(num, "是素数")
			}
		}
	}
	wg3.Done()
}

func main() {
	start := time.Now().Unix()
	// for num := 2; num <= 120000; num++ {
	// 	var flag = true
	// 	for i := 2; i < num; i++ {
	// 		if num%i == 0 {
	// 			flag = false
	// 			break
	// 		}
	// 	}
	// 	if flag {
	// 		fmt.Println(num, "是素数")
	// 	}
	// }

	for i := 1; i <= 4; i++ {
		wg3.Add(1)
		go test2(i)
	}
	wg3.Wait()
	fmt.Println("执行结束")
	end := time.Now().Unix()
	fmt.Println(end - start)
}
