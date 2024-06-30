package MinCost2ReachInTime_1928

import (
	"container/heap"
	"math"
)

// pair 表示边的目标节点和权重
type pair struct {
	first  int
	second int
}

// node 表示优先队列中的节点，包括当前成本、距离和节点编号
type node struct {
	cost int
	dist int
	u    int
}

// priorityQueue 是优先队列类型，用于实现最小堆
type priorityQueue []node

// 以下是优先队列的必要方法

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

// minCost 计算从起点到终点的最小成本
func minCost(maxTime int, edges [][]int, passingFees []int) int {
	n := len(passingFees)
	graph := buildGraph(n, edges)

	return dijkstra(graph, 0, n-1, maxTime, passingFees)
}

// buildGraph 构建图数据结构
func buildGraph(n int, edges [][]int) [][]pair {
	graph := make([][]pair, n)
	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		graph[u] = append(graph[u], pair{v, w})
		graph[v] = append(graph[v], pair{u, w})
	}
	return graph
}

// dijkstra 使用Dijkstra算法计算最短路径的最小成本
func dijkstra(graph [][]pair, src, dst, maxTime int, passingFees []int) int {
	// 初始化成本和距离数组
	cost := make([]int, len(graph))
	dist := make([]int, len(graph))
	for i := range cost {
		cost[i] = math.MaxInt32
	}
	for i := range dist {
		dist[i] = maxTime + 1
	}

	// 使用最小堆优先队列
	pq := &priorityQueue{}
	cost[src] = passingFees[src]
	dist[src] = 0
	heap.Push(pq, node{cost[src], dist[src], src})

	// 进行Dijkstra算法迭代
	for pq.Len() > 0 {
		curr := heap.Pop(pq).(node)
		currCost, d, u := curr.cost, curr.dist, curr.u

		// 如果到达目标节点，返回最小成本
		if u == dst {
			return cost[dst]
		}

		// 如果当前节点的距离大于记录的距离或者当前成本大于记录的成本，则跳过
		if d > dist[u] && currCost > cost[u] {
			continue
		}

		// 遍历当前节点的邻居节点
		for _, edge := range graph[u] {
			v, w := edge.first, edge.second

			// 如果经过当前节点到邻居节点超过最大时间限制，则跳过
			if d+w > maxTime {
				continue
			}

			// 更新邻居节点的成本和距离
			updateCostAndDist(currCost, d, v, w, &cost, &dist, passingFees, pq)
		}
	}

	// 如果无法到达终点节点，则返回-1
	return -1
}

// updateCostAndDist 更新节点的成本和距离，并将更新后的节点推入优先队列
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
