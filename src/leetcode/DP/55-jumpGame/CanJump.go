package _5_jumpGame

func canJump(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[0] = true
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && nums[j] >= i-j {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(nums)-1]
}
