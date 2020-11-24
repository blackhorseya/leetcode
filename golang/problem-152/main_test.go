package main

import "testing"

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
			name: "[2,3,-2,4] then 6",
			args: args{[]int{2, 3, -2, 4}},
			want: 6,
		},
		{
			name: "[-2,0,-1] then 0",
			args: args{[]int{-2, 0, -1}},
			want: 0,
		},
		{
			name: "[-2] then -2",
			args: args{[]int{-2}},
			want: -2,
		},
		{
			name: "[-3,-1,-1] then 3",
			args: args{[]int{-3, -1, -1}},
			want: 3,
		},
		{
			name: "[-4,-3] then 12",
			args: args{[]int{-4, -3}},
			want: 12,
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
