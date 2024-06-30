package MaxNumOfPoints_1937

import "math"

func maxPoints(points [][]int) int64 {
	m := len(points)
	n := len(points[0])
	dp := make([]int64, n)

	// 初始化
	for c := 0; c < n; c++ {
		dp[c] = int64(points[0][c])
	}

	for r := 1; r < m; r++ {
		// 从左到右
		left := make([]int64, n)
		left[0] = dp[0]
		for c := 1; c < n; c++ {
			left[c] = int64(math.Max(float64(left[c-1])-1, float64(dp[c])))
		}

		// 从右到左
		right := make([]int64, n)
		right[n-1] = dp[n-1]
		for c := n - 2; c >= 0; c-- {
			right[c] = int64(math.Max(float64(right[c+1])-1, float64(dp[c])))
		}

		// 更新
		for c := 0; c < n; c++ {
			dp[c] = int64(points[r][c]) + int64(math.Max(float64(left[c]), float64(right[c])))
		}
	}

	maxScore := dp[0]
	for c := 1; c < n; c++ {
		if dp[c] > maxScore {
			maxScore = dp[c]
		}
	}
	return maxScore
}
