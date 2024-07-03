package LargestPailndromeProduct_479

import "math"

func largestPalindrome(n int) int {
	if n == 1 {
		return 9
	}
	num1 := int(math.Pow10(n)) - 1 // 如果是n位数字，最大为10^n-1，举例子如果n是3，则最大的两个数为999*999
	// 这里从回文这个点出发，999*999，n位数相乘，得到的数最大是2n位，所以我们可以确定最终结果的左边n位，然后去翻转得到右边n位，然后计算得到最终结果回文数
	// 得到该回文数后，判断是否可以由n位数相乘得到即可。
	for left := num1; ; left-- {
		n := left
		for r := left; r > 0; r /= 10 {
			n = n*10 + r%10 // 得到回文数结果
		}
		for i := num1; i*i >= n; i-- {
			if n%i == 0 {
				return n % 1337
			}
		}
	}
}

// 打表
func largestPalindrome1(n int) int {
	mp := []int{9, 987, 123, 597, 677, 1218, 877, 475, 475}
	return mp[n-1]
}
