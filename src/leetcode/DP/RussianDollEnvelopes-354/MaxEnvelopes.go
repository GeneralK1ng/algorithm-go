package RussianDollEnvelopes_354

import "sort"

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		return envelopes[j][0] > envelopes[i][0] || a[0] == b[0] && a[1] > b[1]
	})
	var b []int
	for i := 0; i < len(envelopes); i++ {
		p := binarySearch(b, envelopes[i][1])
		if p >= len(b) {
			b = append(b, envelopes[i][1])
		} else {

			b[p] = envelopes[i][1]
		}
	}

	return len(b)
}

func binarySearch(b []int, v int) int {
	l, r := 0, len(b)
	for l < r {
		m := (l + r) / 2
		if b[m] < v {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
