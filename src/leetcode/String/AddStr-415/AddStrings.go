package AddStr_415

import (
	"math/big"
	"strconv"
)

func addStrings1(num1 string, num2 string) string {
	n1 := new(big.Int)
	n2 := new(big.Int)
	n1, _ = n1.SetString(num1, 10)
	n2, _ = n2.SetString(num2, 10)

	sum := new(big.Int).Add(n1, n2)

	return sum.String()
}

func addStrings2(num1 string, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	carry := 0
	result := ""

	for i >= 0 || j >= 0 || carry != 0 {
		var n1, n2 int
		if i >= 0 {
			n1 = int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			n2 = int(num2[j] - '0')
			j--
		}

		sum := n1 + n2 + carry
		carry = sum / 10
		result = strconv.Itoa(sum%10+'0') + result
	}

	return result
}
