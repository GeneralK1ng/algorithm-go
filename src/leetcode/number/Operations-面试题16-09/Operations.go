package Operations_面试题16_09

type Operations struct {
}

func Constructor() Operations {
	return Operations{}
}

func (this *Operations) Minus(a int, b int) int {
	return a - b
}

func (this *Operations) Multiply(a int, b int) int {
	return a * b
}

func (this *Operations) Divide(a int, b int) int {
	return a / b
}
