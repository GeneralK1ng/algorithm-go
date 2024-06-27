package MinStringBalanced_1653

import "testing"

func Test_minimumDeletions1(t *testing.T) {
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
			name: "Test1",
			args: args{
				s: "aababbab",
			},
			want: 2,
		},
		{
			name: "Test2",
			args: args{
				s: "bbaaaaabb",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDeletions1(tt.args.s); got != tt.want {
				t.Errorf("minimumDeletions1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimumDeletions2(t *testing.T) {
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
			name: "Test1",
			args: args{
				s: "aababbab",
			},
			want: 2,
		},
		{
			name: "Test2",
			args: args{
				s: "bbaaaaabb",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDeletions2(tt.args.s); got != tt.want {
				t.Errorf("minimumDeletions2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimumDeletions3(t *testing.T) {
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
			name: "Test1",
			args: args{
				s: "aababbab",
			},
			want: 2,
		},
		{
			name: "Test2",
			args: args{
				s: "bbaaaaabb",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDeletions3(tt.args.s); got != tt.want {
				t.Errorf("minimumDeletions3() = %v, want %v", got, tt.want)
			}
		})
	}
}
