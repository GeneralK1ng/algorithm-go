package BaseCalculate_227

import (
	"strconv"
)

func calculate(s string) int {
	s = removeSpaces(s)
	charArray := []rune(s)
	num, sign, result := 0, 1, 0
	stack := make([]int, 0)

	for i := 0; i < len(charArray); i++ {
		if charArray[i] >= '0' && charArray[i] <= '9' {
			// 处理数字
			num = num*10 + int(charArray[i]-'0')
		} else if charArray[i] == '+' {
			// 处理加号
			result += sign * num
			num, sign = 0, 1
		} else if charArray[i] == '-' {
			// 处理减号
			result += sign * num
			num, sign = 0, -1
		} else if charArray[i] == '*' {
			// 处理乘号
			j := i + 1
			for j < len(charArray) && (charArray[j] == ' ' || (charArray[j] >= '0' && charArray[j] <= '9')) {
				j++
			}
			nextNum, _ := strconv.Atoi(string(charArray[i+1 : j]))
			num *= nextNum
			i = j - 1
		} else if charArray[i] == '/' {
			// 处理除号
			j := i + 1
			for j < len(charArray) && (charArray[j] == ' ' || (charArray[j] >= '0' && charArray[j] <= '9')) {
				j++
			}
			nextNum, _ := strconv.Atoi(string(charArray[i+1 : j]))
			num /= nextNum
			i = j - 1
		} else if charArray[i] == '(' {
			// 处理左括号，保存当前结果和符号，重置num和sign
			stack = append(stack, result, sign)
			result, num, sign = 0, 0, 1
		} else if charArray[i] == ')' {
			// 处理右括号，计算括号内的结果
			result += sign * num
			num = result
			result, sign = stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
		}
	}

	return result + sign*num
}

func removeSpaces(s string) string {
	result := make([]rune, 0, len(s))
	for _, char := range s {
		if char != ' ' {
			result = append(result, char)
		}
	}
	return string(result)
}
