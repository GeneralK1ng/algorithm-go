package CountOfRangeSum_327

func countRangeSum(nums []int, lower, upper int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// 计算前缀和数组
	sums := make([]int, n+1)
	for i := 0; i < n; i++ {
		sums[i+1] = sums[i] + nums[i]
	}

	// 统计符合条件的区间和个数
	var count func(left, right int) int
	count = func(left, right int) int {
		if left >= right {
			return 0
		}
		mid := (left + right) / 2
		countLeft := count(left, mid)
		countRight := count(mid+1, right)
		total := countLeft + countRight

		// 计算跨越左右两部分的区间和个数
		slow, fast := mid+1, mid+1
		for i := left; i <= mid; i++ {
			for slow <= right && sums[slow]-sums[i] < lower {
				slow++
			}
			for fast <= right && sums[fast]-sums[i] <= upper {
				fast++
			}
			total += fast - slow
		}

		// 合并两个已排序的部分
		merged := make([]int, 0, right-left+1)
		i, j := left, mid+1
		for i <= mid && j <= right {
			if sums[i] <= sums[j] {
				merged = append(merged, sums[i])
				i++
			} else {
				merged = append(merged, sums[j])
				j++
			}
		}
		for i <= mid {
			merged = append(merged, sums[i])
			i++
		}
		for j <= right {
			merged = append(merged, sums[j])
			j++
		}
		copy(sums[left:right+1], merged)

		return total
	}

	return count(0, n)
}
