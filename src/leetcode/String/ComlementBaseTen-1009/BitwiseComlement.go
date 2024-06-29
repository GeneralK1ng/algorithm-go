package ComlementBaseTen_1009

import (
	"math"
	"strconv"
)

func bitwiseComplement(n int) int {
	if n == 1 {
		return 0
	} else if n == 0 {
		return 1
	}

	binStr := ""
	for n > 0 {
		binStr = strconv.Itoa(n%2) + binStr
		n /= 2
	}

	inverse := getInverseCode(binStr)

	return toDec(inverse)
}

func toDec(binStr string) int {
	dec := 0
	for i := 0; i < len(binStr); i++ {
		dec += int(binStr[i]-'0') * int(math.Pow(2, float64(len(binStr)-i-1)))
	}
	return dec
}

func getInverseCode(binStr string) string {
	inverseCode := ""
	for i := 0; i < len(binStr); i++ {
		if binStr[i] == '0' {
			inverseCode += "1"
		} else {
			inverseCode += "0"
		}
	}
	return inverseCode
}
