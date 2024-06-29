package BraceExpansionII_1096

import "sort"

func braceExpansionII(expression string) []string {
	ops := []rune{}
	tokens := []map[string]struct{}{}

	for i := 0; i < len(expression); i++ {
		c := rune(expression[i])
		if c == ',' {
			for len(ops) > 0 && ops[len(ops)-1] != '{' && (ops[len(ops)-1] == '+' || ops[len(ops)-1] == '*') {
				tokens = cal(ops[len(ops)-1], tokens)
				ops = ops[:len(ops)-1]
			}
			ops = append(ops, '+')
		} else if c == '{' {
			if i > 0 && (isLetter(rune(expression[i-1])) || expression[i-1] == '}') {
				for len(ops) > 0 && ops[len(ops)-1] == '*' {
					tokens = cal(ops[len(ops)-1], tokens)
					ops = ops[:len(ops)-1]
				}
				ops = append(ops, '*')
			}
			ops = append(ops, '{')
		} else if c == '}' {
			for ops[len(ops)-1] != '{' {
				tokens = cal(ops[len(ops)-1], tokens)
				ops = ops[:len(ops)-1]
			}
			ops = ops[:len(ops)-1]
		} else {
			if i > 0 && (isLetter(rune(expression[i-1])) || expression[i-1] == '}') {
				for len(ops) > 0 && ops[len(ops)-1] == '*' {
					tokens = cal(ops[len(ops)-1], tokens)
					ops = ops[:len(ops)-1]
				}
				ops = append(ops, '*')
			}
			set := map[string]struct{}{string(c): {}}
			tokens = append(tokens, set)
		}
	}

	for len(ops) > 0 {
		tokens = cal(ops[len(ops)-1], tokens)
		ops = ops[:len(ops)-1]
	}

	result := []string{}
	for s := range tokens[0] {
		result = append(result, s)
	}
	sort.Strings(result)
	return result
}

func cal(op rune, tokens []map[string]struct{}) []map[string]struct{} {
	if op == '*' {
		right := tokens[len(tokens)-1]
		left := tokens[len(tokens)-2]
		res := map[string]struct{}{}
		for l := range left {
			for r := range right {
				res[l+r] = struct{}{}
			}
		}
		tokens = tokens[:len(tokens)-2]
		tokens = append(tokens, res)
	} else if op == '+' {
		right := tokens[len(tokens)-1]
		left := tokens[len(tokens)-2]
		for r := range right {
			left[r] = struct{}{}
		}
		tokens = tokens[:len(tokens)-1]
	}
	return tokens
}

func isLetter(c rune) bool {
	return c >= 'a' && c <= 'z'
}
