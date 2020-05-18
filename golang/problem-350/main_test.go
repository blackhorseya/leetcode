package main

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[1,2,2,1] [2,2] return [2,2]",
			args: args{[]int{1, 2, 2, 1}, []int{2, 2}},
			want: []int{2, 2},
		},
		{
			name: "[4,9,5] [9,4,9,8,4] return [4,9]",
			args: args{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}},
			want: []int{4, 9},
		},
		{
			name: "[2,1] [1,2] return [1,2]",
			args: args{[]int{2, 1}, []int{1, 2}},
			want: []int{1, 2},
		},
		{
			name: "[1,2] [1,1] return [1]",
			args: args{[]int{1, 2}, []int{1, 1}},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
