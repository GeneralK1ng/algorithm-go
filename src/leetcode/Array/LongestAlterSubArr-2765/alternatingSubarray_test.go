package LongestAlterSubArr_2765

import "testing"

func Test_alternatingSubArray(t *testing.T) {
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
				nums: []int{2, 3, 4, 3, 4},
			},
			want: 4,
		},
		{
			name: "test2",
			args: args{
				nums: []int{4, 5, 6},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alternatingSubArray(tt.args.nums); got != tt.want {
				t.Errorf("alternatingSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
