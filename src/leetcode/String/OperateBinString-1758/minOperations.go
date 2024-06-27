package OperateBinString_1758

import "math"

func minOperations(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		if int(s[i]-'0') != i%2 {
			result++
		}
	}
	return int(math.Min(float64(result), float64(len(s)-result)))
}
