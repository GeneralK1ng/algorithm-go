package SimplifyPath_71

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
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
				path: "/home/",
			},
			want: "/home",
		},
		{
			name: "Test2",
			args: args{
				path: "/../",
			},
			want: "/",
		},
		{
			name: "Test3",
			args: args{
				path: "/home//foo/",
			},
			want: "/home/foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_simplifyPath2(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		// TODO: Add test cases.
		{
			name: "Test1",
			args: args{
				path: "/home/",
			},
			want: "/home",
		},
		{
			name: "Test2",
			args: args{
				path: "/../",
			},
			want: "/",
		},
		{
			name: "Test3",
			args: args{
				path: "/home//foo/",
			},
			want: "/home/foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath2(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath2() = %v, want %v", got, tt.want)
			}
		})
	}
}
