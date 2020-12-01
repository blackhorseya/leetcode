package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		blocks []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1,5,5,2,6] then 4",
			args: args{[]int{1, 5, 5, 2, 6}},
			want: 4,
		},
		{
			name: "[2,6,8,5] then 3",
			args: args{[]int{2, 6, 8, 5}},
			want: 3,
		},
		{
			name: "[1,1] then 2",
			args: args{[]int{1, 1}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.blocks); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
