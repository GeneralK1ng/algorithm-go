package RemoveBackZeros_2710

import "regexp"

func removeTrailingZeros1(num string) string {
	// 定义一个正则表达式，用于匹配尾随零
	re := regexp.MustCompile(`0+$`)
	// 使用正则表达式替换尾随零为空字符串
	result := re.ReplaceAllString(num, "")
	return result
}

func removeTrailingZeros2(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		if num[i] != '0' {
			return num[:i+1]
		}
	}
	return ""
}
