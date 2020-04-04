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
			name: "123_return_321",
			args: args{
				x: 123,
			},
			want: 321,
		},
		{
			name: "-123_return_-321",
			args: args{
				x: -123,
			},
			want: -321,
		},
		{
			name: "120_return_21",
			args: args{
				x: 120,
			},
			want: 21,
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