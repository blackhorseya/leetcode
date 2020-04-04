package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1,1,2]_return_2",
			args: args{
				nums: []int{1,1,2},
			},
			want: 2,
		},
		{
			name: "[0,0,1,1,1,2,2,3,3,4]_return_5",
			args: args{
				nums: []int{0,0,1,1,1,2,2,3,3,4},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.nums); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}