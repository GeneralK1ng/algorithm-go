package ReplacingTime_1736

import "testing"

func Test_maximumTime(t *testing.T) {
	type args struct {
		time string
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
				time: "2?:?0",
			},
			want: "23:50",
		},
		{
			name: "test2",
			args: args{
				time: "0?:3?",
			},
			want: "09:39",
		},
		{
			name: "test3",
			args: args{
				time: "1?:22",
			},
			want: "19:22",
		},
		{
			name: "test4",
			args: args{
				time: "??:3?",
			},
			want: "23:39",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumTime(tt.args.time); got != tt.want {
				t.Errorf("maximumTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
