package SmallestRange_632

import "container/heap"

// Pair 结构体表示堆中的元素，包含值、列表索引和在列表中的索引
type Pair struct {
	value, listIdx, numIdx int
}

// PairHeap 定义为 Pair 的最小堆
type PairHeap []Pair

func (h PairHeap) Len() int           { return len(h) }
func (h PairHeap) Less(i, j int) bool { return h[i].value < h[j].value }
func (h PairHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PairHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *PairHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *PairHeap) Peek() Pair {
	return (*h)[0]
}

func smallestRange(nums [][]int) []int {
	h := &PairHeap{}
	heap.Init(h)

	// 初始化堆，将每个列表的第一个元素入堆，并记录每个列表的最大值
	maxValues := make([]int, len(nums))
	for i, list := range nums {
		heap.Push(h, Pair{list[0], i, 0})
		maxValues[i] = list[0]
	}

	var (
		minStart, maxEnd = h.Peek().value, getMax(maxValues)
		result           = []int{minStart, maxEnd}
	)

	for h.Len() == len(nums) {
		curr := heap.Pop(h).(Pair)

		// 更新最小区间
		if maxEnd-curr.value < result[1]-result[0] || (maxEnd-curr.value == result[1]-result[0] && curr.value < result[0]) {
			minStart, maxEnd = curr.value, getMax(maxValues)
			result = []int{minStart, maxEnd}
		}

		// 如果当前弹出元素不是其所在列表的最后一个元素，则将下一个元素入堆
		if curr.numIdx+1 < len(nums[curr.listIdx]) {
			heap.Push(h, Pair{nums[curr.listIdx][curr.numIdx+1], curr.listIdx, curr.numIdx + 1})
			maxValues[curr.listIdx] = nums[curr.listIdx][curr.numIdx+1]
		} else {
			// 已经遍历完某个列表，直接结束
			break
		}
	}

	return result
}

func getMax(arr []int) int {
	maxValue := arr[0]
	for _, v := range arr[1:] {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}
