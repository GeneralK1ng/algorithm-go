package DefuseTheBomb_1652

func decrypt(code []int, k int) []int {
	n := len(code)
	result := make([]int, n)

	for i := 0; i < n; i++ {
		sum := 0

		if k > 0 {
			for j := 1; j <= k; j++ {
				sum += code[(i+j)%n]
			}
		} else if k < 0 {
			for j := -1; j >= k; j-- {
				sum += code[(i+j+n)%n]
			}
		}

		result[i] = sum
	}

	return result
}
