package _19_Pascal_s_Triangle_II

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}
	res := make([]int, rowIndex+1)
	res[0] = 1
	res[1] = 1
	for i := 2; i <= rowIndex; i++ {
		for j := i; j >= 1; j-- {
			res[j] = res[j] + res[j-1]
		}
	}
	return res
}
