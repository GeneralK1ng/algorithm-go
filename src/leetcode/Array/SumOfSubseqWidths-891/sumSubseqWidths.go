package SumOfSubseqWidths_891

import "sort"

func sumOfWidths(nums []int) int {
	MOD := 1000000007
	n := len(nums)

	sort.Ints(nums)

	// 初始化总宽度
	totalWidth := 0
	// 初始化存储 2 的幂次的数组
	pow := make([]int, n)
	pow[0] = 1

	// 通过位运算计算 2 的幂次
	for i := 1; i < n; i++ {
		pow[i] = (pow[i-1] << 1) % MOD
	}

	for i := 0; i < n; i++ {
		contrib := (nums[i] * (pow[i] - pow[n-i-1] + MOD)) % MOD
		totalWidth = (totalWidth + contrib) % MOD
	}

	return totalWidth
}
