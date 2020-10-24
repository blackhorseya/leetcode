package main

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		rowIndex int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "0 then [1]",
			args: args{0},
			want: []int{1},
		},
		{
			name: "1 then [1,1]",
			args: args{1},
			want: []int{1, 1},
		},
		{
			name: "2 then [1,2,1]",
			args: args{2},
			want: []int{1, 2, 1},
		},
		{
			name: "3 then [1,3,3,1]",
			args: args{3},
			want: []int{1, 3, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.rowIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
