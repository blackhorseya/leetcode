package main

import (
	"leetcode/base"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		root *base.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nil return 0",
			args: args{nil},
			want: 0,
		},
		{
			name: "[] return 0",
			args: args{&base.TreeNode{}},
			want: 0,
		},
		{
			name: "[1] return 1",
			args: args{&base.TreeNode{Val: 1}},
			want: 0,
		},
		{
			name: "[1,2,3,4,5] return 3",
			args: args{&base.TreeNode{
				Val:   1,
				Left:  &base.TreeNode{Val: 2, Left: &base.TreeNode{Val: 4}, Right: &base.TreeNode{Val: 5}},
				Right: &base.TreeNode{Val: 3},
			}},
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
