package RussianDollEnvelopes_354

import "testing"

func Test_maxEnvelopes(t *testing.T) {
	type args struct {
		envelopes [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}}, 3},
		{"case2", args{[][]int{{1, 1}, {1, 1}, {1, 1}}}, 1},
		{"case3", args{[][]int{{4, 5}, {4, 6}, {6, 7}, {2, 3}, {1, 1}}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEnvelopes(tt.args.envelopes); got != tt.want {
				t.Errorf("maxEnvelopes() = %v, want %v", got, tt.want)
			}
		})
	}
}
