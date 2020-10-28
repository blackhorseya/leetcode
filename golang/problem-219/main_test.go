package main

import "testing"

func TestSolve(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "[1,2,3,1,2,3] 2 then false",
			args: args{[]int{1, 2, 3, 1, 2, 3}, 2},
			want: false,
		},
		{
			name: "[1,0,1,1] 1 then true",
			args: args{[]int{1, 0, 1, 1}, 1},
			want: true,
		},
		{
			name: "[1,2,3,1] 3 then true",
			args: args{[]int{1, 2, 3, 1}, 3},
			want: true,
		},
		{
			name: "[99,99] 2 then true",
			args: args{[]int{99, 99}, 2},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
