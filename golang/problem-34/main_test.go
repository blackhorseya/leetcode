package main

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[] 0 then [-1,-1]",
			args: args{[]int{}, 0},
			want: []int{-1, -1},
		},
		{
			name: "[5,7,7,8,8,10] 8 then [3,4]",
			args: args{[]int{5, 7, 7, 8, 8, 10}, 8},
			want: []int{3, 4},
		},
		{
			name: "[5,7,7,8,8,10] 6 then [-1,-1]",
			args: args{[]int{5, 7, 7, 8, 8, 10}, 6},
			want: []int{-1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
