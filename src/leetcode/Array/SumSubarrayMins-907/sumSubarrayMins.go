package SumSubarrayMins_907

func sumSubarrayMins(arr []int) int {
	n := len(arr)
	left, right := make([]int, n), make([]int, n)
	stackLeft, stackRight := make([]int, 0), make([]int, 0)

	// 计算每个元素的左边界
	for i := 0; i < n; i++ {
		for len(stackLeft) > 0 && arr[i] < arr[stackLeft[len(stackLeft)-1]] {
			stackLeft = stackLeft[:len(stackLeft)-1]
		}
		if len(stackLeft) == 0 {
			left[i] = i + 1
		} else {
			left[i] = i - stackLeft[len(stackLeft)-1]
		}
		stackLeft = append(stackLeft, i)
	}

	// 清空栈
	stackLeft = stackLeft[:0]

	// 计算每个元素的右边界
	for i := n - 1; i >= 0; i-- {
		for len(stackRight) > 0 && arr[i] <= arr[stackRight[len(stackRight)-1]] {
			stackRight = stackRight[:len(stackRight)-1]
		}
		if len(stackRight) == 0 {
			right[i] = n - i
		} else {
			right[i] = stackRight[len(stackRight)-1] - i
		}
		stackRight = append(stackRight, i)
	}

	result, mod := 0, 1000000007

	// 计算以每个元素为最小值的子数组的和
	for i := 0; i < n; i++ {
		result = (result + arr[i]*left[i]*right[i]) % mod
	}

	return result
}
