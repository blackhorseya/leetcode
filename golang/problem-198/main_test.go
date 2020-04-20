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
			name: "nil return 0",
			args: args{
				nums: nil,
			},
			want: 0,
		},
		{
			name: "[1] return 1",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "[1,2] return 2",
			args: args{
				nums: []int{1, 2},
			},
			want: 2,
		},
		{
			name: "[1,2,3] return 4",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: 4,
		},
		{
			name: "[1,2,3,1] return 4",
			args: args{
				nums: []int{1, 2, 3, 1},
			},
			want: 4,
		},
		{
			name: "[2,7,9,3,1] return 12",
			args: args{
				nums: []int{2, 7, 9, 3, 1},
			},
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
