package kthSmallestDistance_719

import "sort"

func kthSmallestDistance(nums []int, k int) int {
	sort.Ints(nums)
	low, high := 0, nums[len(nums)-1]-nums[0]

	for low < high {
		mid := (low + high) / 2
		count := countPairs(nums, mid)

		if count < k {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}

func countPairs(nums []int, target int) int {
	count := 0
	left := 0

	for right := 0; right < len(nums); right++ {
		for left < right && nums[right]-nums[left] > target {
			left++
		}
		count += right - left
	}

	return count
}
