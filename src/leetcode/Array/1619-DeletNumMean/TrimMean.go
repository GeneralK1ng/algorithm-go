package _619_DeletNumMean

import "sort"

func trimMean(arr []int) float64 {
	sort.Ints(arr)
	sum := 0
	for i := int(float64(len(arr)) * 0.05); i < len(arr)-int(float64(len(arr))*0.05); i++ {
		sum += arr[i]
	}

	return float64(sum) / float64(len(arr)-int(float64(len(arr))*0.1))
}
