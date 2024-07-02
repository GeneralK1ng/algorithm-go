package MaxPrimeDist_3115

var primeNumbers = []int{
	2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61,
	67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137,
	139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211,
	227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293,
	293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379}

// 判断一个数是否是质数
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for _, prime := range primeNumbers {
		if prime == num {
			return true
		}
		if num%prime == 0 {
			return false
		}
	}
	return true
}

// 找出数组中质数的下标
func findPrimeIndices(nums []int) []int {
	var primeIndices []int
	for i, num := range nums {
		if isPrime(num) {
			primeIndices = append(primeIndices, i)
		}
	}
	return primeIndices
}

// 计算最大距离
func maximumPrimeDifference(nums []int) int {
	var m, n int = -1, -1
	for i := 0; i < len(nums); i++ {
		if isPrime(nums[i]) {
			if m == -1 {
				m = i
			} else {
				n = i
			}
		}
	}

	if n == -1 {
		return 0
	}
	return n - m
}
