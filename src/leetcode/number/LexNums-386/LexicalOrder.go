package LexNums_386

func lexicalOrder(n int) []int {
	res := make([]int, 0, n)
	for i := 1; i <= 9 && i <= n; i++ {
		res = append(res, i)
		dfs(i, n, &res)
	}
	return res
}

func dfs(cur, n int, res *[]int) {
	if len(*res) == n {
		return
	}
	for i := 0; i < 10; i++ {
		next := cur*10 + i
		if next > n {
			break
		}
		*res = append(*res, next)
		dfs(next, n, res)
	}
}
