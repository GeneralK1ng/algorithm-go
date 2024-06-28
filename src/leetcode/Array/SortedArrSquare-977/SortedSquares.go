package SortedArrSquare_977

import "sort"

func sortedSquares(nums []int) []int {
	res := squareArr(nums)

	sort.Ints(res)
	return res
}

func squareArr(nums []int) []int {
	for i := range nums {
		nums[i] = nums[i] * nums[i]
	}
	return nums
}
