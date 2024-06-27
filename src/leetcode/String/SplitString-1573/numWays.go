package SplitString_1573

func numWays(s string) int {
	mod := 1000000007
	ones := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			ones++
		}
	}

	if ones%3 != 0 {
		return 0
	}

	if ones == 0 {
		return (len(s) - 1) * (len(s) - 2) / 2 % mod
	}

	N, a, b, c, d, count := ones/3, 0, 0, 0, 0, 0
	for i, letter := range s {
		if letter == '0' {
			continue
		}
		if letter == '1' {
			count++
		}
		if count == N {
			a = i
		}
		if count == N+1 {
			b = i
		}
		if count == 2*N {
			c = i
		}
		if count == 2*N+1 {
			d = i
		}
	}
	return (b - a) * (d - c) % mod

}
