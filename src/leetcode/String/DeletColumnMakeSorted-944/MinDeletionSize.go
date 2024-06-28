package DeletColumnMakeSorted_944

func minDeletionSize(strs []string) int {
	if len(strs) == 0 || len(strs[0]) == 0 {
		return 0
	}

	cnt := 0
	rows, cols := len(strs), len(strs[0])

	// 只判断列
	for col := 0; col < cols; col++ {
		for row := 1; row < rows; row++ {
			if strs[row][col] < strs[row-1][col] {
				cnt++
				break
			}
		}
	}
	return cnt
}
