package TargetSum_494

func findTargetSumWays(nums []int, target int) int {
	n := len(nums)

	sum := 0
	for _, num := range nums {
		sum += num
	}

	if target > sum || target < -sum {
		return 0
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 2*sum+1)
	}

	dp[0][sum] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= 2*sum; j++ {
			if j >= nums[i-1] {
				dp[i][j] += dp[i-1][j-nums[i-1]]
			}
			if j+nums[i-1] <= 2*sum {
				dp[i][j] += dp[i-1][j+nums[i-1]]
			}
		}
	}

	return dp[n][target+sum]
}
