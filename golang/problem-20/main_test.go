package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "()_return_true",
			args: args{
				s: "()",
			},
			want: true,
		},
		{
			name: "()[]{}_return_true",
			args: args{
				s: "()[]{}",
			},
			want: true,
		},
		{
			name: "([)]_return_false",
			args: args{
				s: "([)]",
			},
			want: false,
		},
		{
			name: "(]_return_false",
			args: args{
				s: "(]",
			},
			want: false,
		},
		{
			name: "{[]}_return_true",
			args: args{
				s: "{[]}",
			},
			want: true,
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