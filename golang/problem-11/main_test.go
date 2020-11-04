package main

import "testing"

func TestSolve(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1,8,6,2,5,4,8,3,7] then 49",
			args: args{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
			want: 49,
		},
		{
			name: "[1,1] then 1",
			args: args{[]int{1, 1}},
			want: 1,
		},
		{
			name: "[1,2,1] then 2",
			args: args{[]int{1, 2, 1}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.height); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
