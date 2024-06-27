package kthSmallestDistance_719

import "testing"

func Test_kthSmallestDistance(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				nums: []int{1, 3, 1},
				k:    1,
			},
			want: 0,
		},
		{
			name: "test 2",
			args: args{
				nums: []int{1, 1, 1},
				k:    2,
			},
			want: 0,
		},
		{
			name: "test 3",
			args: args{
				nums: []int{1, 6, 1},
				k:    3,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallestDistance(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("kthSmallestDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
