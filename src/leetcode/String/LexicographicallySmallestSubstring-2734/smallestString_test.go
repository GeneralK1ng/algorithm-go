package LexicographicallySmallestSubstring_2734

import "testing"

func Test_smallestString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"1", args{"cbabc"}, "baabc"},
		{"2", args{"aa"}, "az"},
		{"3", args{"acbbc"}, "abaab"},
		{"4", args{"leetcode"}, "kddsbncd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestString(tt.args.s); got != tt.want {
				t.Errorf("smallestString() = %v, want %v", got, tt.want)
			}
		})
	}
}
