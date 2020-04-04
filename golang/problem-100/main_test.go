package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		p *TreeNode
		q *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "[1,2,3]_[1,2,3]_return_true",
			args: args{
				p: &TreeNode{
					Val:   1,
					Left:  &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
				q: &TreeNode{
					Val:   1,
					Left:  &TreeNode{
						Val:   2,
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
			want: true,
		},
		{
			name: "[1,2]_[1,null,2]_return_false",
			args: args{
				p: &TreeNode{
					Val:   1,
					Left:  &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
				},
				q: &TreeNode{
					Val:   1,
					Right: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
				},
			},
			want: false,
		},
		{
			name: "[1,2,1]_[1,1,2]_return_false",
			args: args{
				p: &TreeNode{
					Val:   1,
					Left:  &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   1,
						Left:  nil,
						Right: nil,
					},
				},
				q: &TreeNode{
					Val:   1,
					Left:  &TreeNode{
						Val:   1,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}