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
		want bool
	}{
		{
			name: "121_return_true",
			args: args{
				x: 121,
			},
			want: true,
		},
		{
			name: "-121_return_false",
			args: args{
				x: -121,
			},
			want: false,
		},
		{
			name: "10_return_false",
			args: args{
				x: 10,
			},
			want: false,
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