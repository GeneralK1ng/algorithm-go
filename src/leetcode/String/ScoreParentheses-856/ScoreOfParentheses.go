package ScoreParentheses_856

func scoreOfParentheses(s string) int {
	// 单调栈
	var stack []int
	stack = append(stack, 0)
	for _, v := range s {
		if v == '(' {
			stack = append(stack, 0)
		} else {
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] += max(tmp*2, 1)
		}
	}
	return stack[len(stack)-1]
}
