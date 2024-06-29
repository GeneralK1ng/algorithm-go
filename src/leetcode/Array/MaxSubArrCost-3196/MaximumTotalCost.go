package MaxSubArrCost_3196

import "math"

func maximumTotalCost(nums []int) int64 {
	n := len(nums)
	var dp = make([][2]int64, n)

	dp[0][0] = int64(nums[0])
	dp[0][1] = math.MinInt64

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]) + int64(nums[i])
		dp[i][1] = dp[i-1][0] - int64(nums[i])
	}

	return max(dp[n-1][0], dp[n-1][1])
}
