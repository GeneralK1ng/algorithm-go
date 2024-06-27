package CountWords_2085

import "testing"

func TestCountWords(t *testing.T) {
	type args struct {
		words1 []string
		words2 []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				words1: []string{"leetcode", "is", "amazing", "as", "is"},
				words2: []string{"amazing", "leetcode", "is"},
			},
			want: 2,
		},
		{
			name: "test2",
			args: args{
				words1: []string{"b", "bb", "bbb"},
				words2: []string{"a", "aa", "aaa"},
			},
			want: 0,
		},
		{
			name: "test3",
			args: args{
				words1: []string{"a", "ab"},
				words2: []string{"a", "a", "a", "ab"},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countWords(tt.args.words1, tt.args.words2); got != tt.want {
				t.Errorf("countWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
