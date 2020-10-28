package main

import (
	"leetcode/base"
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		t1 *base.TreeNode
		t2 *base.TreeNode
	}
	tests := []struct {
		name string
		args args
		want *base.TreeNode
	}{
		{
			name: "[1] nil return [1]",
			args: args{&base.TreeNode{Val: 1}, nil},
			want: &base.TreeNode{Val: 1},
		},
		{
			name: "nil [1] return [1]",
			args: args{nil, &base.TreeNode{Val: 1}},
			want: &base.TreeNode{Val: 1},
		},
		{
			name: "[1] [1] return [2]",
			args: args{&base.TreeNode{Val: 1}, &base.TreeNode{Val: 1}},
			want: &base.TreeNode{Val: 2},
		},
		{
			name: "[1,2,3] [1,2,3] return [2,4,6]",
			args: args{
				&base.TreeNode{Val: 1, Left: &base.TreeNode{Val: 2}, Right: &base.TreeNode{Val: 3}},
				&base.TreeNode{Val: 1, Left: &base.TreeNode{Val: 2}, Right: &base.TreeNode{Val: 3}}},
			want: &base.TreeNode{Val: 2, Left: &base.TreeNode{Val: 4}, Right: &base.TreeNode{Val: 6}},
		},
		{
			name: "[1,nil,3] [1,2,3] return [2,4,6]",
			args: args{
				&base.TreeNode{Val: 1, Left: nil, Right: &base.TreeNode{Val: 3}},
				&base.TreeNode{Val: 1, Left: &base.TreeNode{Val: 2}, Right: &base.TreeNode{Val: 3}}},
			want: &base.TreeNode{Val: 2, Left: &base.TreeNode{Val: 2}, Right: &base.TreeNode{Val: 6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.t1, tt.args.t2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
