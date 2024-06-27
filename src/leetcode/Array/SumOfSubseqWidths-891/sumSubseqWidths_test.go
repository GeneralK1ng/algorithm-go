package SumOfSubseqWidths_891

import "testing"

func Test_sumOfWidths(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				nums: []int{2, 1, 3},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfWidths(tt.args.nums); got != tt.want {
				t.Errorf("sumOfWidths() = %v, want %v", got, tt.want)
			}
		})
	}
}
