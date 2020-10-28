package main

import "testing"

func TestSolve(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "[1,2,3,4] then false",
			args: args{[]int{1, 2, 3, 4}},
			want: false,
		},
		{
			name: "[1,1,1,3,3,4,3,2,4,2] then true",
			args: args{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}},
			want: true,
		},
		{
			name: "[1,2,3,1] then true",
			args: args{[]int{1, 2, 3, 1}},
			want: true,
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
