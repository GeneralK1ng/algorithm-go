package SmallGoodBase_483

import (
	"math"
	"strconv"
)

func smallestGoodBase(n string) string {
	// 将字符串形式的n转换为整数
	num, _ := strconv.ParseInt(n, 10, 64)

	// 最大可能的位数m，因为k^(digit+1)必定大于n
	digit := int(math.Log2(float64(num)))

	for bits := digit; bits > 0; bits-- {
		// k^digit = n + 1
		k := int64(math.Pow(float64(num+1), 1/float64(bits)))

		var sum, power int64 = 1, 1
		for i := 0; i < bits; i++ {
			power *= k
			sum += power
		}
		if sum == num {
			return strconv.FormatInt(k, 10)
		}
	}
	return strconv.FormatInt(num-1, 10)
}
