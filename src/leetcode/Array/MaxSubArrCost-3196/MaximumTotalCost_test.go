package MaxSubArrCost_3196

import "testing"

func Test_maximumTotalCost(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"1", args{[]int{1, -2, 3, 4}}, 10},
		{"2", args{[]int{1, -1, 1, -1}}, 4},
		{"3", args{[]int{0}}, 0},
		{"4", args{[]int{1, -1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumTotalCost(tt.args.nums); got != tt.want {
				t.Errorf("maximumTotalCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
