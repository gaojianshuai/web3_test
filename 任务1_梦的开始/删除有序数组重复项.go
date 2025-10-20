
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	number := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[number] {
			number++
			nums[number] = nums[i]
		}
	}
	return number + 1
}