package TribonacciNum_1137

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}

	dp := make([]int, n+2)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}
	return dp[n]
}
