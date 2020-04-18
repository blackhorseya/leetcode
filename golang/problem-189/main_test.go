package main

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[1,2] 3 return [2,1]",
			args: args{
				nums: []int{1, 2},
				k:    3,
			},
			want: []int{2, 1},
		},
		{
			name: "[1,2,3,4,5,6,7] 3 return [5,6,7,1,2,3,4]",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "[-1,-100,3,99] 2 return [3,99,-1,-100]",
			args: args{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
			want: []int{3, 99, -1, -100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Solve(tt.args.nums, tt.args.k); !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("Solve() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
