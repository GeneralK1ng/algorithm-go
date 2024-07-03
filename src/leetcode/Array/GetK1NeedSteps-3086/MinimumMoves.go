package GetK1NeedSteps_3086

import "math"

// 计算前缀和数组
func calculatePrefixSums(nums []int) ([]int64, []int64) {
	n := len(nums)
	indexSum := make([]int64, n+1)
	sum := make([]int64, n+1)
	for i := 0; i < n; i++ {
		indexSum[i+1] = indexSum[i] + int64(nums[i]*i)
		sum[i+1] = sum[i] + int64(nums[i])
	}
	return indexSum, sum
}

// 计算当前位置的邻居之和
func calculateNeighborSum(i int, nums []int) int {
	neighborSum := nums[i]
	if i-1 >= 0 {
		neighborSum += nums[i-1]
	}
	if i+1 < len(nums) {
		neighborSum += nums[i+1]
	}
	return neighborSum
}

// 计算在给定步数范围内是否能收集到k个1
func canCollectKOnes(nums []int, sum []int64, k, maxChanges, steps int) bool {
	n := len(nums)
	for i := 0; i <= n-steps; i++ {
		if sum[i+steps]-sum[i] >= int64(k-maxChanges) {
			return true
		}
	}
	return false
}

// 使用二分查找来决定最小步数
func findMinSteps(nums []int, sum []int64, k, maxChanges int) int64 {
	n := len(nums)
	left, right := 0, n
	for left < right {
		mid := (left + right) / 2
		if canCollectKOnes(nums, sum, k, maxChanges, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return int64(left)
}

// 主函数
func minimumMoves(nums []int, k int, maxChanges int) int64 {
	n := len(nums)
	indexSum, sum := calculatePrefixSums(nums)
	minMoves := int64(math.MaxInt64)

	for i := 0; i < n; i++ {
		neighborSum := calculateNeighborSum(i, nums)
		if neighborSum+maxChanges >= k {
			if k <= neighborSum {
				minMoves = int64(math.Min(float64(minMoves), float64(k-nums[i])))
			} else {
				minMoves = int64(math.Min(float64(minMoves), float64(2*k-neighborSum-nums[i])))
			}
			continue
		}

		left, right := 0, n
		for left <= right {
			mid := (left + right) / 2
			i1 := int(math.Max(float64(i-mid), 0))
			i2 := int(math.Min(float64(i+mid), float64(n-1)))
			if sum[i2+1]-sum[i1] >= int64(k-maxChanges) {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		i1 := int(math.Max(float64(i-left), 0))
		i2 := int(math.Min(float64(i+left), float64(n-1)))
		if sum[i2+1]-sum[i1] > int64(k-maxChanges) {
			i1++
		}
		count1 := sum[i+1] - sum[i1]
		count2 := sum[i2+1] - sum[i+1]
		minMoves = int64(math.Min(float64(minMoves), float64(indexSum[i2+1]-indexSum[i+1]-int64(i)*count2+int64(i)*count1-(indexSum[i+1]-indexSum[i1])+2*int64(maxChanges))))
	}
	return minMoves
}
