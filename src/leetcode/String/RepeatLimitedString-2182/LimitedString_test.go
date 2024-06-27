package RepeatLimitedString_2182

import "testing"

func TestRepeatLimitedString(t *testing.T) {
	type args struct {
		input       string
		repeatLimit int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Test1",
			args: args{
				input:       "cczazcc",
				repeatLimit: 3,
			},
			want: "zzcccac",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatLimitedString(tt.args.input, tt.args.repeatLimit); got != tt.want {
				t.Errorf("repeatLimitedString() = %v, want %v", got, tt.want)
			}
		})
	}
}
