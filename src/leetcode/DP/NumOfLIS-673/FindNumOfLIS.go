package NumOfLIS_673

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	dp := make([]int, n)
	count := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = 1
		count[i] = 1
	}

	maxLength := 1
	result := 1

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					count[i] = count[j]
				} else if dp[j]+1 == dp[i] {
					count[i] += count[j]
				}
			}
		}
		if dp[i] > maxLength {
			maxLength = dp[i]
			result = count[i]
		} else if dp[i] == maxLength {
			result += count[i]
		}
	}

	return result
}
