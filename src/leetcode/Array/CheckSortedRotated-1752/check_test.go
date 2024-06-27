package CheckSortedRotated_1752

import "testing"

func Test_check(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				nums: []int{3, 4, 5, 1, 2},
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				nums: []int{2, 1, 3, 4},
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := check(tt.args.nums); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}
