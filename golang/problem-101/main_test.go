package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "[1,2,2,3,4,4,3]_return_true",
			args: args{
				root: &TreeNode{
					Val:   1,
					Left:  &TreeNode{
						Val:   2,
						Left:  &TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   4,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val:   2,
						Left:  &TreeNode{
							Val:   4,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			want: true,
		},
		{
			name: "[1,2,2,null,3,null,3]_return_false",
			args: args{
				root: &TreeNode{
					Val:   1,
					Left:  &TreeNode{
						Val:   2,
						Left:  nil,
						Right: &TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: &TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
					},
				},},
			want: false,
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