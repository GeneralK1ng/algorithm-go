package String2Fibonacci_842

import (
	"reflect"
	"testing"
)

func Test_splitIntoFibonacci(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "Test1",
			args: args{
				num: "1101111",
			},
			want: []int{11, 0, 11, 11},
		},
		{
			name: "Test2",
			args: args{
				num: "112358130",
			},
			want: []int{},
		},
		{
			name: "Test3",
			args: args{
				num: "0123",
			},
			want: []int{},
		},
		{
			name: "Test4",
			args: args{
				num: "17522",
			},
			want: []int{17, 5, 22},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitIntoFibonacci(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitIntoFibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
