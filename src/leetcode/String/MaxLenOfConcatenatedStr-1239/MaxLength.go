package MaxLenOfConcatenatedStr_1239

import (
	"math/bits"
)

func maxLength(arr []string) int {
	var (
		ans   int
		masks []int
		dfs   func(int, int)
	)
outer:
	for _, s := range arr {
		mask := 0
		for _, ch := range s {
			ch -= 'a'
			if mask>>ch&1 > 0 {
				continue outer
			}
			mask |= 1 << ch
		}
		masks = append(masks, mask)
	}

	n := len(masks)
	dfs = func(i int, mask int) {
		if i == n {
			ans = max(ans, bits.OnesCount(uint(mask)))
			return
		}
		if mask&masks[i] == 0 {
			dfs(i+1, mask|masks[i])
		}
		dfs(i+1, mask)
	}
	dfs(0, 0)
	return ans
}
