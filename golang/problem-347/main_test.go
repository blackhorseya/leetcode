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
			name: "[1,1,1,2,2,3] 2 then [1,2]",
			args: args{[]int{1, 1, 1, 2, 2, 3}, 2},
			want: []int{1, 2},
		},
		{
			name: "[1] 1 then [1]",
			args: args{[]int{1}, 1},
			want: []int{1},
		},
		{
			name: "[1,2] 2 then [1,2]",
			args: args{[]int{1, 2}, 2},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
