package LargestAltitude_1732

func largestAltitude(gain []int) int {
	var arr = make([]int, len(gain)*2)

	// 初始化第一个元素
	arr[0] = 0

	// 循环计算该数组的前序和
	for i := range gain {
		arr[i+1] = arr[i] + gain[i]
	}

	max := 0
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}

func largestAltitudePro(gain []int) int {

	max, curAltitude := 0, 0

	for _, g := range gain {
		curAltitude += g
		if curAltitude > max {
			max = curAltitude
		}
	}

	return max
}
