package String2Fibonacci_842

import (
	"math"
	"strconv"
)

func splitIntoFibonacci(num string) []int {
	var result []int
	backtrack(&result, num, 0)
	return result
}

func backtrack(result *[]int, num string, start int) bool {
	if start == len(num) && len(*result) >= 3 {
		return true
	}

	for i := start; i < len(num); i++ {
		currentStr := num[start : i+1]

		if (currentStr[0] == '0' && len(currentStr) > 1) || toInt(currentStr) > math.MaxInt32 {
			continue
		}

		currentNum := toInt(currentStr)

		if len(*result) >= 2 {
			target := (*result)[len(*result)-1] + (*result)[len(*result)-2]
			if currentNum > target {
				break
			}
			if currentNum < target {
				continue
			}
		}

		*result = append(*result, currentNum)

		if backtrack(result, num, i+1) {
			return true
		}

		*result = (*result)[:len(*result)-1]
	}
	return false
}

func toInt(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}
