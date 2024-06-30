package CntNumOfSquareSubset_2572

const mod = 1_000_000_007

func squareFreeSubsets(nums []int) int {
	count := make([]int, 31)
	for _, num := range nums {
		count[num]++
	}

	result := int64(0)
	result = int64(count[2]+1) * int64(count[3]+1) * int64(count[5]+1) * int64(count[7]+1) * int64(count[11]+1) * int64(count[13]+1) * int64(count[17]+1) * int64(count[19]+1) * int64(count[23]+1) * int64(count[29]+1) % mod
	result += int64(count[6]) * int64(count[5]+1) * int64(count[7]+1) * int64(count[11]+1) * int64(count[13]+1) * int64(count[17]+1) * int64(count[19]+1) * int64(count[23]+1) * int64(count[29]+1) % mod
	result += int64(count[10]) * int64(count[3]+1) * int64(count[7]+1) * int64(count[11]+1) * int64(count[13]+1) * int64(count[17]+1) * int64(count[19]+1) * int64(count[23]+1) * int64(count[29]+1) % mod
	result += int64(count[14]) * int64(count[3]+1) * int64(count[5]+1) * int64(count[11]+1) * int64(count[13]+1) * int64(count[17]+1) * int64(count[19]+1) * int64(count[23]+1) * int64(count[29]+1) % mod
	result += int64(count[22]) * int64(count[3]+1) * int64(count[5]+1) * int64(count[7]+1) * int64(count[13]+1) * int64(count[17]+1) * int64(count[19]+1) * int64(count[23]+1) * int64(count[29]+1) % mod
	result += int64(count[26]) * int64(count[3]+1) * int64(count[5]+1) * int64(count[7]+1) * int64(count[11]+1) * int64(count[17]+1) * int64(count[19]+1) * int64(count[23]+1) * int64(count[29]+1) % mod
	result += int64(count[2]+1) * int64(count[15]) * int64(count[7]+1) * int64(count[11]+1) * int64(count[13]+1) * int64(count[17]+1) * int64(count[19]+1) * int64(count[23]+1) * int64(count[29]+1) % mod
	result += int64(count[2]+1) * int64(count[21]) * int64(count[5]+1) * int64(count[11]+1) * int64(count[13]+1) * int64(count[17]+1) * int64(count[19]+1) * int64(count[23]+1) * int64(count[29]+1) % mod
	result += int64(count[10]) * int64(count[21]) * int64(count[11]+1) * int64(count[13]+1) * int64(count[17]+1) * int64(count[19]+1) * int64(count[23]+1) * int64(count[29]+1) % mod
	result += int64(count[14]) * int64(count[15]) * int64(count[11]+1) * int64(count[13]+1) * int64(count[17]+1) * int64(count[19]+1) * int64(count[23]+1) * int64(count[29]+1) % mod
	result += int64(count[22]) * int64(count[15]) * int64(count[7]+1) * int64(count[13]+1) * int64(count[17]+1) * int64(count[19]+1) * int64(count[23]+1) * int64(count[29]+1) % mod
	result += int64(count[22]) * int64(count[21]) * int64(count[5]+1) * int64(count[13]+1) * int64(count[17]+1) * int64(count[19]+1) * int64(count[23]+1) * int64(count[29]+1) % mod
	result += int64(count[26]) * int64(count[15]) * int64(count[7]+1) * int64(count[11]+1) * int64(count[17]+1) * int64(count[19]+1) * int64(count[23]+1) * int64(count[29]+1) % mod
	result += int64(count[26]) * int64(count[21]) * int64(count[5]+1) * int64(count[11]+1) * int64(count[17]+1) * int64(count[19]+1) * int64(count[23]+1) * int64(count[29]+1) % mod
	result += int64(count[30]) * int64(count[7]+1) * int64(count[11]+1) * int64(count[13]+1) * int64(count[17]+1) * int64(count[19]+1) * int64(count[23]+1) * int64(count[29]+1) % mod

	for i := 0; i < count[1]; i++ {
		result = (result * 2) % mod
	}

	result = (result - 1 + mod) % mod

	return int(result)
}
