package MinOperation2Arr_2208

import "container/heap"

func halveArray(nums []int) int {
	total := 0.0

	for _, num := range nums {
		total += float64(num)
	}

	target := total / 2

	h := &MaxHeap{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, float64(num))
	}

	currentSum := total
	operations := 0

	for currentSum > target {
		maxValue := heap.Pop(h).(float64)
		halfValue := maxValue / 2
		currentSum -= halfValue
		heap.Push(h, halfValue)
		operations++
	}

	return operations
}

type MaxHeap []float64

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(float64))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
