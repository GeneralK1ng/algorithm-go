package MaxIncreasingGroups_2790

import (
	"sort"
)

func maxIncreasingGroups(usageLimits []int) int {
	sort.Ints(usageLimits)
	now := usageLimits[0] - 1
	res := 1

	for _, value := range usageLimits[1:] {
		now += value
		if now >= res+1 {
			now -= res + 1
			res += 1
		}
	}

	return res
}
