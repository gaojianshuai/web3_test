//考察：数组操作、进位处理

// 题目：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	result := make([]int, len(digits)+1)
	result[0] = 1
	return result
}