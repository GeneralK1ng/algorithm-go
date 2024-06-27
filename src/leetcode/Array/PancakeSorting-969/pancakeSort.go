package PancakeSorting_969

func pancakeSort(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	right := len(arr)
	var (
		ans []int
	)
	for right > 0 {
		idx := find(arr, right)
		if idx != right-1 {
			reversePancake(arr, 0, idx)
			reversePancake(arr, 0, right-1)
			ans = append(ans, idx+1, right)
		}
		right--
	}

	return ans
}

func reversePancake(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func find(nums []int, right int) int {
	for i, num := range nums {
		if num == right {
			return i
		}
	}
	return -1
}
