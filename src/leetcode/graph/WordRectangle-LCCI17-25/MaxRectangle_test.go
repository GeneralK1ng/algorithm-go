package WordRectangle_LCCI17_25

import (
	"reflect"
	"testing"
)

func Test_maxRectangle(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{[]string{"this", "real", "hard", "trh", "hea", "iar", "sld"}},
			[]string{"this", "real", "hard"}},
		{"2", args{[]string{"aa"}}, []string{"aa", "aa"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxRectangle(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
