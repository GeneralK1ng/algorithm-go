package MinStringBalanced_1653

import "strings"

func minimumDeletions1(s string) int {
	a, b, res := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			a++
		}
	}
	res = a
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			a--
		} else {
			b++
		}
		res = min(res, a+b)
	}
	return res
}

func minimumDeletions2(s string) int {
	prev, res, b := 0, 0, 0
	for _, c := range s {
		if c == 'a' {
			res = min(prev+1, b)
			prev = res
		} else {
			b++
		}
	}
	return res
}

func minimumDeletions3(s string) int {
	del := strings.Count(s, "a")
	mini := del

	for _, c := range s {
		del += int((c-'a')*2 - 1)
		mini = min(mini, del)
	}
	return mini
}
