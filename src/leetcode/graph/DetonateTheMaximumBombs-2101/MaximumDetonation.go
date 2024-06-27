package DetonateTheMaximumBombs_2101

import (
	"math"
	"sort"
)

type Bomb struct {
	x, y, r int
}

func maximumDetonation(bombs [][]int) int {
	n := len(bombs)
	bombList := make([]Bomb, n)
	for i := 0; i < n; i++ {
		bombList[i] = Bomb{bombs[i][0], bombs[i][1], bombs[i][2]}
	}

	// 按照爆炸范围从大到小排序
	sort.Slice(bombList, func(i, j int) bool {
		return bombList[i].r > bombList[j].r
	})

	maxBombsDetonated := 1
	visited := make([]bool, n)
	var queue []int

	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		queue = append(queue, i)
		curVisited := make([]bool, n)
		curVisited[i] = true
		visited[i] = true
		count := 1

		for len(queue) > 0 {
			idx := queue[0]
			queue = queue[1:]
			for j := 0; j < n; j++ {
				if !curVisited[j] && withinRange(bombList[idx], bombList[j]) {
					queue = append(queue, j)
					curVisited[j] = true
					visited[j] = true
					count++
				}
			}
		}
		if count > maxBombsDetonated {
			maxBombsDetonated = count
		}
	}

	return maxBombsDetonated
}

func withinRange(b1, b2 Bomb) bool {
	distanceSquared := math.Pow(float64(b2.x-b1.x), 2) + math.Pow(float64(b2.y-b1.y), 2)
	return distanceSquared <= math.Pow(float64(b1.r), 2)
}
