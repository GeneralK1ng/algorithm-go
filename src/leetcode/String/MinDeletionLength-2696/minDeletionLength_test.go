package MinDeletionLength_2696

import "testing"

func Test_minLength(t *testing.T) {
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
				s: "ABFCACDB",
			},
			want: 2,
		},
		{
			name: "Test2",
			args: args{
				s: "ACBBD",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minLength(tt.args.s); got != tt.want {
				t.Errorf("minLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minLengthByStack(t *testing.T) {
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
				s: "ABFCACDB",
			},
			want: 2,
		},
		{
			name: "Test2",
			args: args{
				s: "ACBBD",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minLengthByStack(tt.args.s); got != tt.want {
				t.Errorf("minLengthByStack() = %v, want %v", got, tt.want)
			}
		})
	}
}
