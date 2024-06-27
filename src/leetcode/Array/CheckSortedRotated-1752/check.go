package CheckSortedRotated_1752

func check(nums []int) bool {
	n := len(nums)
	count := 0

	// 遍历数组
	for i := 0; i < n; i++ {
		// 如果当前元素大于下一个元素
		if nums[i] > nums[(i+1)%n] {
			count++
		}

		// 如果逆序对超过1个，返回false
		if count > 1 {
			return false
		}
	}

	// 如果逆序对不超过1个，返回true
	return true
}
