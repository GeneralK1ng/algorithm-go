package MinCost2ReachInTime_1928

import (
	"container/heap"
	"math"
)

type pair struct {
	first  int
	second int
}

type node struct {
	cost int
	dist int
	u    int
}

type priorityQueue []node

func (pq priorityQueue) Len() int           { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool { return pq[i].cost < pq[j].cost }
func (pq priorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *priorityQueue) Push(x interface{}) {
	item := x.(node)
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func minCost(maxTime int, edges [][]int, passingFees []int) int {
	n := len(passingFees)
	graph := buildGraph(n, edges)

	return dijkstra(graph, 0, n-1, maxTime, passingFees)
}

func buildGraph(n int, edges [][]int) [][]pair {
	graph := make([][]pair, n)
	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		graph[u] = append(graph[u], pair{v, w})
		graph[v] = append(graph[v], pair{u, w})
	}
	return graph
}

func dijkstra(graph [][]pair, src, dst, maxTime int, passingFees []int) int {
	cost := make([]int, len(graph))
	dist := make([]int, len(graph))
	for i := range cost {
		cost[i] = math.MaxInt32
	}
	for i := range dist {
		dist[i] = maxTime + 1
	}

	pq := &priorityQueue{}
	cost[src] = passingFees[src]
	dist[src] = 0
	heap.Push(pq, node{cost[src], dist[src], src})

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(node)
		currCost, d, u := curr.cost, curr.dist, curr.u

		if u == dst {
			return cost[dst]
		}

		if d > dist[u] && currCost > cost[u] {
			continue
		}

		for _, edge := range graph[u] {
			v, w := edge.first, edge.second

			if d+w > maxTime {
				continue
			}

			updateCostAndDist(currCost, d, v, w, &cost, &dist, passingFees, pq)
		}
	}

	return -1
}

func updateCostAndDist(currCost, d, v, w int, cost, dist *[]int, passingFees []int, pq *priorityQueue) {
	if currCost+passingFees[v] < (*cost)[v] {
		(*cost)[v] = currCost + passingFees[v]
		(*dist)[v] = d + w
		heap.Push(pq, node{(*cost)[v], (*dist)[v], v})
	} else if d+w < (*dist)[v] {
		(*dist)[v] = d + w
		heap.Push(pq, node{currCost + passingFees[v], (*dist)[v], v})
	}
}
