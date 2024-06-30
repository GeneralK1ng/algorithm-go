package LeastNumOfKRemove_1481

import "sort"

// findLeastNumOfUniqueInts 函数用于找出在给定数组中，通过移除最少数量的元素后，剩余元素中最少的不重复整数的数量。
// arr 是一个整数数组，k 是一个整数，表示可以移除的元素总数。
// 返回值是在移除元素后，剩余元素中最少的不重复整数的数量。
func findLeastNumOfUniqueInts(arr []int, k int) int {
	// 使用 map 统计每个元素出现的次数。
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}

	// 将每个元素出现的次数收集到一个切片中。
	var res []int
	for _, v := range m {
		res = append(res, v)
	}

	// 对出现次数进行排序，以便从最少出现次数的元素开始移除。
	sort.Ints(res)

	// 遍历排序后的出现次数，尝试移除元素。
	for i := 0; i < len(res); i++ {
		// 如果当前可用的移除次数大于等于当前元素的出现次数，移除该次数，并更新可用移除次数。
		if k >= res[i] {
			k -= res[i]
			res[i] = 0
		} else {
			// 如果当前可用的移除次数不足以移除当前元素，返回还需要移除的不重复元素数量。
			return len(res) - i
		}
	}

	// 如果所有元素都可以被完全移除，返回 0。
	return 0
}
