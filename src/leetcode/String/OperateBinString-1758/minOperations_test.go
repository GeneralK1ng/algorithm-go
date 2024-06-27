package OperateBinString_1758

import "testing"

func Test_minOperations(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Test_1",
			args: args{
				s: "0100",
			},
			want: 1,
		},
		{
			name: "Test_2",
			args: args{
				s: "10",
			},
			want: 0,
		},
		{
			name: "Test_3",
			args: args{
				s: "1111",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.s); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
