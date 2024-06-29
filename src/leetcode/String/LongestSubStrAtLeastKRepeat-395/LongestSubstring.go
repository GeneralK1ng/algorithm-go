package LongestSubStrAtLeastKRepeat_395

func longestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}
	return findLongestSubstring(s, k)
}

func findLongestSubstring(s string, k int) int {
	freq := make(map[rune]int)
	for _, char := range s {
		freq[char]++
	}

	for i, char := range s {
		if freq[char] < k {
			left := findLongestSubstring(s[:i], k)
			right := findLongestSubstring(s[i+1:], k)

			if left > right {
				return left
			}
			return right
		}
	}

	return len(s)
}
