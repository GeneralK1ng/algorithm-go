package MinAddToMakeValid_2645

import "testing"

func Test_addMinimum(t *testing.T) {
	type args struct {
		word string
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
				word: "b",
			},
			want: 2,
		},
		{
			name: "test2",
			args: args{
				word: "aaa",
			},
			want: 6,
		},
		{
			name: "test3",
			args: args{
				word: "ab",
			},
			want: 1,
		},
		{
			name: "test4",
			args: args{
				word: "abc",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addMinimum(tt.args.word); got != tt.want {
				t.Errorf("addMinimum() = %v, want %v", got, tt.want)
			}
		})
	}
}
