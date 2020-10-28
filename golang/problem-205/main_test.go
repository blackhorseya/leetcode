package main

import "testing"

func TestSolve(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "foo bar then false",
			args: args{"foo", "bar"},
			want: false,
		},
		{
			name: "egg add then true",
			args: args{"egg", "add"},
			want: true,
		},
		{
			name: "title paper then true",
			args: args{"title", "paper"},
			want: true,
		},
		{
			name: "ab aa then false",
			args: args{"ab", "aa"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
