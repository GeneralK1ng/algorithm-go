package WordPuzzle_LCR129

var directions = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func wordPuzzle(grid [][]byte, target string) bool {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	var dfs func(int, int, int) bool
	dfs = func(i, j, k int) bool {
		if k == len(target) { // 找到完整单词
			return true
		}
		if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] || grid[i][j] != target[k] {
			return false
		}
		visited[i][j] = true // 标记已访问
		for _, dir := range directions {
			x, y := i+dir[0], j+dir[1]
			if dfs(x, y, k+1) {
				return true
			}
		}
		visited[i][j] = false // 回溯，撤销标记
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
