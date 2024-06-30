package RobotPrintSmallStr_2434

import "testing"

func Test_robotWithString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"zza"}, "azz"},
		{"2", args{"bac"}, "abc"},
		{"3", args{"bdda"}, "addb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := robotWithString(tt.args.s); got != tt.want {
				t.Errorf("robotWithString() = %v, want %v", got, tt.want)
			}
		})
	}
}
