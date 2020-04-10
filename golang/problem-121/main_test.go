package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nil_return_0",
			args: args{
				prices: nil,
			},
			want: 0,
		},
		{
			name: "[]_return_0",
			args: args{
				prices: []int{},
			},
			want: 0,
		},
		{
			name: "[1]_return_1",
			args: args{
				prices: []int{1},
			},
			want: 0,
		},
		{
			name: "[1,3]_return_2",
			args: args{
				prices: []int{1, 3},
			},
			want: 2,
		},
		{
			name: "[7,1,5,3,6,4]_return_5",
			args: args{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			want: 5,
		},
		{
			name: "[7,6,4,3,1]_return_0",
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.prices); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
