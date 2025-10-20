func twoSum(nums []int, target int) []int {
	// var arr9 = [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				fmt.Println(i, j)
				return []int{i, j}
			}
		}
	}
	return nil // 没有找到
}
