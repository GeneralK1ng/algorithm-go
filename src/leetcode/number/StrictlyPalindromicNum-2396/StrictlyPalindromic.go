package StrictlyPalindromicNum_2396

func isStrictlyPalindromic(n int) bool {
	for i := 2; i < n-1; i++ {
		if !isPalindrome(n, i) {
			return false
		}
	}
	return true
}

func isPalindrome(n, base int) bool {
	var res int
	for n > 0 {
		res = res*base + n%base
		n /= base
	}
	return res == n
}
