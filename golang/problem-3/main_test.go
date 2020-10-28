package main

import "testing"

func TestSolve(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a return 1",
			args: args{"a"},
			want: 1,
		},
		{
			name: "bbbbb return 1",
			args: args{"bbbbb"},
			want: 1,
		},
		{
			name: "abcabcbb return 3",
			args: args{"abcabcbb"},
			want: 3,
		},
		{
			name: "pwwkew return 3",
			args: args{"pwwkew"},
			want: 3,
		},
		{
			name: "au return 2",
			args: args{"au"},
			want: 2,
		},
		{
			name: "dvdf return 3",
			args: args{"dvdf"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.s); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
