package SummaryRange_228

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	var result []string
	n := len(nums)

	addRange := func(start, end int) {
		if start == end {
			result = append(result, strconv.Itoa(start))
		} else {
			result = append(result, fmt.Sprintf("%d->%d", start, end))
		}
	}

	if n == 0 {
		return result
	}

	start := nums[0]

	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1]+1 {
			addRange(start, nums[i-1])
			start = nums[i]
		}
	}

	addRange(start, nums[n-1])

	return result
}
