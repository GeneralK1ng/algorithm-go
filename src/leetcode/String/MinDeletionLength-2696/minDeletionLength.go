package MinDeletionLength_2696

import "strings"

func minLength(s string) int {
	// 哈希表用于缓存计算结果
	cache := make(map[string]int)

	// 递归函数，删除所有 "AB" 和 "CD" 子串
	var removeABCD func(string) int
	removeABCD = func(str string) int {
		// 如果字符串已经在缓存中存在，直接返回缓存结果
		if val, ok := cache[str]; ok {
			return val
		}

		newStr := strings.Replace(str, "AB", "", -1)
		newStr = strings.Replace(newStr, "CD", "", -1)

		// 如果字符串未发生变化，说明不能再删除子串，返回当前字符串长度
		if newStr == str {
			return len(newStr)
		}

		// 递归调用删除子串操作，并将结果缓存
		result := removeABCD(newStr)
		cache[str] = result
		return result
	}

	// 调用递归函数，并返回最终字符串的长度
	result := removeABCD(s)
	return result
}

func minLengthByStack(s string) int {
	var stack []byte // 用栈来模拟删除子串的操作

	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && (stack[len(stack)-1] == 'A' && s[i] == 'B' || stack[len(stack)-1] == 'C' && s[i] == 'D') {
			// 如果当前字符与栈顶字符构成"AB"或"CD"子串，执行删除操作
			stack = stack[:len(stack)-1]
		} else {
			// 否则，将当前字符入栈
			stack = append(stack, s[i])
		}
	}

	// 最终栈中的字符即为可获得的最终字符串
	return len(stack)
}
