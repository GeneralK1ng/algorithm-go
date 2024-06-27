package DiverseSubstrings_1748B

import (
	"fmt"
)

func solve() {
	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)

	ans := 0
	for i := 0; i < n; i++ {
		cnt := make([]int, 10)
		uniqueCount, maxFrequency := 0, 0

		for j := i; j < min(n, i+100); j++ {
			u := int(s[j] - '0')
			cnt[u]++
			if cnt[u] == 1 {
				uniqueCount++
			}
			maxFrequency = max(maxFrequency, cnt[u])
		}

		if maxFrequency <= uniqueCount {
			ans++
		}
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var t int
	fmt.Scan(&t)
	for ; t > 0; t-- {
		solve()
	}
}
