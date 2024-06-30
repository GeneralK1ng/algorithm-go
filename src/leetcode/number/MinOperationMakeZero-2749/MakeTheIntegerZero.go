package MinOperationMakeZero_2749

import "math/bits"

func makeTheIntegerZero(num1, num2 int) int {
	for i := 1; ; i++ {
		x := num1 - i*num2
		if x < 0 {
			break
		}
		if bits.OnesCount(uint(x)) <= i && i <= x {
			return i
		}
	}
	return -1
}
