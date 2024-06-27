package ReorganizeString_767

func reorganizeString(s string) string {
	n := len(s)
	count := make([]int, 26)
	maxCount, maxIndex := 0, 0

	// 统计每个字符的出现次数，并找到出现次数最多的字符
	for _, c := range s {
		count[c-'a']++
		if count[c-'a'] > maxCount {
			maxCount = count[c-'a']
			maxIndex = int(c - 'a')
		}
	}

	if maxCount > (n+1)/2 {
		return ""
	}

	result := make([]byte, n)
	index := 0

	// 先填充出现次数最多的字符
	for count[maxIndex] > 0 {
		result[index] = byte(maxIndex) + 'a'
		index += 2
		count[maxIndex]--
	}

	// 填充其他字符
	for i := 0; i < 26; i++ {
		for count[i] > 0 {
			if index >= n {
				index = 1
			}
			result[index] = byte(i) + 'a'
			index += 2
			count[i]--
		}
	}

	return string(result)
}
