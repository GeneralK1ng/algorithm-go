package RepeatLimitedString_2182

func repeatLimitedString(s string, repeatLimit int) string {
	// 统计每个字母出现的次数
	letterCount := make([]int, 26)
	length := len(s)

	// 如果字符串长度为1，直接返回
	if length == 1 {
		return s
	}

	// 统计字母表中字母数量，O(n)
	for i := 0; i < length; i++ {
		letterCount[s[i]-'a']++
	}

	currentLetter, nextLetter := 25, 1
	count := 0
	result := make([]byte, length)

	// 重新构造，O(n)
	for i := 0; i < length; i++ {
		// 达到重复限制，需要交换字母
		// 当前字母数量大于0，且达到重复限制
		if count >= repeatLimit && letterCount[currentLetter] > 0 {
			// 找下一个字符用来做分割
			nextLetter = 1
			for {
				if nextLetter > currentLetter {
					// 没有可用来分割的字符，退出
					return string(result[:i])
				} else if letterCount[currentLetter-nextLetter] == 0 {
					nextLetter++
				} else {
					break
				}
			}
			result[i] = byte(currentLetter - nextLetter + 'a')
			letterCount[currentLetter-nextLetter]--
			count = 0 // 分割，重置计数
		} else {
			// 当前字母数量大于0，直接使用
			if letterCount[currentLetter] > 0 {
				result[i] = byte(currentLetter + 'a')
				letterCount[currentLetter]--
				count++
			} else {
				// 移动到下一个非零字母桶
				for letterCount[currentLetter] == 0 {
					currentLetter--
				}
				result[i] = byte(currentLetter + 'a')
				letterCount[currentLetter]--
				count = 1 // 重置计数
			}
		}
	}
	return string(result)
}
