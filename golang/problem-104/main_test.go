package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	var tests = []struct {
		name string
		args args
		want int
	}{
		{
			name: "[3,9,20,null,null,15,7]_return_3",
			args: args{
				root: &TreeNode{
					Val:   3,
					Left:  &TreeNode{
						Val:   9,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   20,
						Left:  &TreeNode{
							Val:   15,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   7,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.root); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}