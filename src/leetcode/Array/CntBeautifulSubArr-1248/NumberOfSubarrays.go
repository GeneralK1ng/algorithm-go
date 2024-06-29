package CntBeautifulSubArr_1248

func numberOfSubarrays(nums []int, k int) int {
	var oddIndices []int
	count := 0

	for i := 0; i < len(nums); i++ {
		if nums[i]%2 != 0 {
			oddIndices = append(oddIndices, i)
		}
	}

	for i := 0; i <= len(oddIndices)-k; i++ {
		count += countSubarrays(nums, oddIndices[i], oddIndices[i+k-1])
	}

	return count
}

func countSubarrays(nums []int, left, right int) int {
	left--
	right++
	leftEvens := 0
	rightEvens := 0
	result := 1

	for left >= 0 {
		if nums[left]%2 == 0 {
			leftEvens++
			left--
		} else {
			break
		}
	}

	for right < len(nums) {
		if nums[right]%2 == 0 {
			rightEvens++
			right++
		} else {
			break
		}
	}

	return result + leftEvens + rightEvens + leftEvens*rightEvens
}
