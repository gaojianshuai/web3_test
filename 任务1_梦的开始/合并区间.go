func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// 按起始位置升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		last := result[len(result)-1]

		// 判断是否重叠
		if current[0] <= last[1] {
			// 合并区间
			newStart := last[0]
			newEnd := max(last[1], current[1])
			result[len(result)-1] = []int{newStart, newEnd}
		} else {
			result = append(result, current)
		}
	}

	return result
}