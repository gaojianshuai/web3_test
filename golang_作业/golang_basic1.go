package main

import "fmt"

func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num // 对每个数字进行异或运算
	}
	return result // 最终result将是只出现一次的数字
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	result := singleNumber(nums)
	fmt.Println("The single number is:", result) // 输出应为4，因为我们知道除了4外，其他数字都出现了两次
}
