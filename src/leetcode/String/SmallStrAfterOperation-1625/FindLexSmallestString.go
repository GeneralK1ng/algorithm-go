package SmallStrAfterOperation_1625

func findLexSmallestString(s string, a int, b int) string {
	n := len(s)
	best := s
	visited := make(map[string]bool)
	stack := []string{s}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[current] {
			continue
		}
		visited[current] = true

		if current < best {
			best = current
		}

		nextAccumulate := make([]byte, n)
		for i := 0; i < n; i++ {
			if i%2 == 1 {
				nextAccumulate[i] = byte((int(current[i]-'0')+a)%10 + '0')
			} else {
				nextAccumulate[i] = current[i]
			}
		}
		nextAccStr := string(nextAccumulate)
		stack = append(stack, nextAccStr)

		nextRotate := rotateString(current, b)
		stack = append(stack, nextRotate)
	}

	return best
}

func rotateString(s string, b int) string {
	n := len(s)
	b %= n
	return s[n-b:] + s[:n-b]
}
