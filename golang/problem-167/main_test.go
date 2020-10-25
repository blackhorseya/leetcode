package main

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[] 1 then []",
			args: args{nil, 1},
			want: nil,
		},
		{
			name: "[-1,0] -1 then [1,2]",
			args: args{[]int{-1, 0}, -1},
			want: []int{1, 2},
		},
		{
			name: "[2,3,4] 6 then [1,3]",
			args: args{[]int{2, 3, 4}, 6},
			want: []int{1, 3},
		},
		{
			name: "[2,7,11,15] 9 then [1,2]",
			args: args{[]int{2, 7, 11, 15}, 9},
			want: []int{1, 2},
		},
		{
			name: "[0,0,3,4] 0 then [1,2]",
			args: args{[]int{0, 0, 3, 4}, 0},
			want: []int{1, 2},
		},
		{
			name: "[0,4,4,10] 8 then [2,3]",
			args: args{[]int{0, 4, 4, 10}, 8},
			want: []int{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
