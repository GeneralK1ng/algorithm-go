package BaseCalculate_227

import "testing"

func Test_calculate(t *testing.T) {
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
			name: "test1",
			args: args{
				s: "3+2*2",
			},
			want: 7,
		},
		{
			name: "test2",
			args: args{
				s: " 3/2 ",
			},
			want: 1,
		},
		{
			name: "test3",
			args: args{
				s: " 3+5 / 2 ",
			},
			want: 5,
		},
		{
			name: "test4",
			args: args{
				s: "(1+(4+5+2)-3) * 14",
			},
			want: 126,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
