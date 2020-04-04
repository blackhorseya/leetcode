package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "4_return_2",
			args: args{
				x: 4,
			},
			want: 2,
		},
		{
			name: "8_return_2",
			args: args{
				x: 8,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.x); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
