package MinimumCostToMergeStones_1000

import (
	"math"
)

func mergeStones(stones []int, k int) int {
	n := len(stones)
	if (n-1)%(k-1) != 0 {
		return -1
	}

	// 前缀和
	prefixSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + stones[i]
	}

	// dp[l][r]表示将stones[l:r+1]合并成尽可能少的堆的最小成本
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}

	// 初始化
	for i := 0; i < n; i++ {
		dp[i][i] = 0
	}

	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			for m := i; m < j; m += k - 1 {
				dp[i][j] = min(dp[i][j], dp[i][m]+dp[m+1][j])
			}
			if (j-i)%(k-1) == 0 {
				dp[i][j] += prefixSum[j+1] - prefixSum[i]
			}
		}
	}

	return dp[0][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
