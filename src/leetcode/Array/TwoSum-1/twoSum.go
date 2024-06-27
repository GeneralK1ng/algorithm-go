package TwoSum_1

func twoSum(nums []int, target int) []int {
	// 创建一个map用于存储数组元素和它们的索引
	numsMap := make(map[int]int)

	// 遍历数组
	for i, num := range nums {
		// 计算当前元素与目标值的差值
		complement := target - num

		// 检查差值是否在map中存在
		if idx, found := numsMap[complement]; found {
			// 找到了两个数的和等于目标值，返回它们的索引
			return []int{idx, i}
		}

		// 将当前元素及其索引存入map
		numsMap[num] = i
	}

	// 没有找到满足条件的两个数，抛出异常
	panic("No two sum solution")
}
