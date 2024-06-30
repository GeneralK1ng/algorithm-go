package FindKClosestElements_658

import "sort"

func findClosestElements(arr []int, k int, x int) []int {
	index := sort.SearchInts(arr, x)
	start := index
	end := index

	for start > 0 && end < len(arr) && end-start < k {
		if x-arr[start-1] <= arr[end]-x {
			start--
		} else {
			end++
		}
	}
	for end-start < k {
		if start == 0 {
			end++
		} else if end == len(arr) {
			start--
		} else {
			if x-arr[start-1] <= arr[end]-x {
				start--
			} else {
				end++
			}
		}
	}

	return arr[start:end]
}
