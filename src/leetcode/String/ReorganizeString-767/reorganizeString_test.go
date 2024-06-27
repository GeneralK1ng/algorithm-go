package ReorganizeString_767

import "testing"

func Test_reorganizeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				s: "aab",
			},
			want: "aba",
		},
		{
			name: "test2",
			args: args{
				s: "aaab",
			},
			want: "",
		},
		{
			name: "test3",
			args: args{
				s: "abc",
			},
			want: "acb",
		},
		{
			name: "test4",
			args: args{
				s: "aaaabbbbcccc",
			},
			want: "ababacacbcbc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorganizeString(tt.args.s); got != tt.want {
				t.Errorf("reorganizeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
