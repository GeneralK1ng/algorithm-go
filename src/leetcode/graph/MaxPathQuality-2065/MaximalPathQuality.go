package MaxPathQuality_2065

func maximalPathQuality(values []int, edges [][]int, maxTime int) int {
	n := len(values)
	g := make([][][]int, n)

	for _, edge := range edges {
		u, v, t := edge[0], edge[1], edge[2]
		g[u] = append(g[u], []int{v, t})
		g[v] = append(g[v], []int{u, t})
	}

	visited := make([]bool, n)
	res := 0

	var dfs func(u, time, value int)
	dfs = func(u, time, value int) {
		if u == 0 {
			res = max(res, value)
		}
		for _, neighbor := range g[u] {
			v := neighbor[0]
			dist := neighbor[1]
			if time-dist >= 0 {
				if !visited[v] {
					visited[v] = true
					dfs(v, time-dist, value+values[v])
					visited[v] = false
				} else {
					dfs(v, time-dist, value)
				}
			}
		}
	}

	visited[0] = true
	dfs(0, maxTime, values[0])

	return res
}
