package main

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "[-10,-3,0,5,9]_return_[0,-3,9,-10,null,5]",
			args: args{
				nums: []int{-10,-3,0,5,9},
			},
			want: &TreeNode{
				Val:   0,
				Left:  &TreeNode{
					Val:   -3,
					Left:  &TreeNode{
						Val:   -10,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
				Right: &TreeNode{
					Val:   9,
					Left:  &TreeNode{
						Val:   5,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}