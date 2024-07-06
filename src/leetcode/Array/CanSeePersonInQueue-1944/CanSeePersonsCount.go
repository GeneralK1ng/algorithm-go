package CanSeePersonInQueue_1944

func canSeePersonsCount(heights []int) []int {
	var stack []int
	result := make([]int, len(heights))
	for i := len(heights) - 1; i >= 0; i-- {

		for len(stack) > 0 && stack[len(stack)-1] <= heights[i] {
			stack = stack[:len(stack)-1]
			result[i] += 1
		}

		if len(stack) > 0 {
			result[i] += 1
		}
		stack = append(stack, heights[i])
	}
	return result
}
