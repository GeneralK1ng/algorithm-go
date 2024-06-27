package MaxIncreasingGroups_2790

import "testing"

func Test_maxIncreasingGroups(t *testing.T) {
	type args struct {
		usageLimits []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			args: args{
				usageLimits: []int{1, 2, 5},
			},
			want: 3,
		},
		{
			args: args{
				usageLimits: []int{2, 1, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxIncreasingGroups(tt.args.usageLimits); got != tt.want {
				t.Errorf("maxIncreasingGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}
