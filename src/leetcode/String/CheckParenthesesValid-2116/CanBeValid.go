package CheckParenthesesValid_2116

func canBeValid(s string, locked string) bool {
	if len(s)%2 != 0 {
		return false
	}

	open := 0
	free := 0

	// 从左到右遍历
	for i := 0; i < len(s); i++ {
		if locked[i] == '1' {
			if s[i] == '(' {
				open++
			} else {
				open--
			}
		} else {
			free++
		}
		if open < -free {
			return false
		}
	}

	close := 0
	free = 0

	// 从右到左遍历
	for i := len(s) - 1; i >= 0; i-- {
		if locked[i] == '1' {
			if s[i] == ')' {
				close++
			} else {
				close--
			}
		} else {
			free++
		}
		if close < -free {
			return false
		}
	}

	return true
}
