package main

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "[1,3],[2,6],[8,10],[15,18] then [1,6],[8,10],[15,18]",
			args: args{[][]int{
				{1, 3},
				{2, 6},
				{8, 10},
				{15, 18},
			}},
			want: [][]int{
				{1, 6},
				{8, 10},
				{15, 18},
			},
		},
		{
			name: "[1,4],[4,5] then [1,5]",
			args: args{[][]int{
				{1, 4},
				{4, 5},
			}},
			want: [][]int{
				{1, 5},
			},
		},
		{
			name: "[1,4],[0,4] then [0,4]",
			args: args{[][]int{
				{1, 4},
				{0, 4},
			}},
			want: [][]int{
				{0, 4},
			},
		},
		{
			name: "[1,4],[0,1] then [0,4]",
			args: args{[][]int{
				{1, 4},
				{0, 1},
			}},
			want: [][]int{
				{0, 4},
			},
		},
		{
			name: "[1,4],[0,2],[3,5] then [0,5]",
			args: args{[][]int{
				{1, 4},
				{0, 2},
				{3, 5},
			}},
			want: [][]int{
				{0, 5},
			},
		},
		{
			name: "[2,3],[4,5],[6,7],[8,9],[1,10] then [1,10]",
			args: args{[][]int{
				{2, 3},
				{4, 5},
				{6, 7},
				{8, 9},
				{1, 10},
			}},
			want: [][]int{
				{1, 10},
			},
		},
		{
			name: "[][] then nil",
			args: args{[][]int{}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_include(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "[1,3] [2,6] then true",
			args: args{[]int{1, 3}, []int{2, 6}},
			want: true,
		},
		{
			name: "[1,4] [0,1] then true",
			args: args{[]int{1, 4}, []int{0, 1}},
			want: true,
		},
		{
			name: "[1,4] [0,4] then true",
			args: args{[]int{1, 4}, []int{0, 4}},
			want: true,
		},
		{
			name: "[1,4] [4,5] then true",
			args: args{[]int{1, 4}, []int{4, 5}},
			want: true,
		},
		{
			name: "[1,4] [0,0] then false",
			args: args{[]int{1, 4}, []int{0, 0}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := include(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("include() = %v, want %v", got, tt.want)
			}
		})
	}
}
