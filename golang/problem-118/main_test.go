package main

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		numRows int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1_return_[][1]",
			args: args{
				numRows: 1,
			},
			want: [][]int{{1}},
		},
		{
			name: "2_return_[[1],[1,1]]",
			args: args{
				numRows: 2,
			},
			want: [][]int{{1},{1,1}},
		},
		{
			name: "3_return_[[1],[1,1],[1,2,1]]",
			args: args{
				numRows: 3,
			},
			want: [][]int{{1},{1,1},{1,2,1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.numRows); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
