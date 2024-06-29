package KeysAndRooms_841

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	visited := make([]bool, n)
	visited[0] = true

	dfs(visited, rooms, 0)

	for _, v := range visited {
		if !v {
			return false
		}
	}

	return true
}

func dfs(visited []bool, rooms [][]int, room int) {
	visited[room] = true
	for _, key := range rooms[room] {
		if !visited[key] {
			dfs(visited, rooms, key)
		}
	}
}
