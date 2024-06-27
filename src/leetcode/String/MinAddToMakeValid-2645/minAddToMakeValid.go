package MinAddToMakeValid_2645

func addMinimum(word string) int {
	// 创建一个缓存用于存储已经计算过的结果
	cache := make(map[string]int)
	return addMinimumHelper(word, cache)
}

func addMinimumHelper(word string, cache map[string]int) int {
	// 如果字符串为空，返回0
	if len(word) == 0 {
		return 0
	}

	// 如果已经计算过该子问题，直接返回缓存中的结果
	if val, ok := cache[word]; ok {
		return val
	}

	// 递归计算并保存结果到缓存
	if len(word) >= 3 && word[:3] == "abc" {
		cache[word] = addMinimumHelper(word[3:], cache)
	} else if len(word) >= 2 && (word[:2] == "ab" || word[:2] == "ac" || word[:2] == "bc") {
		cache[word] = 1 + addMinimumHelper(word[2:], cache)
	} else if word[0] == 'a' || word[0] == 'b' || word[0] == 'c' {
		cache[word] = 2 + addMinimumHelper(word[1:], cache)
	} else {
		cache[word] = 0
	}

	return cache[word]
}
