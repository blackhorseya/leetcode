package main

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[1,2,3,0,0,0]_3_[2,5,6]_3_return_[1,2,2,3,5,6]",
			args: args{
				nums1: []int{1,2,3,0,0,0},
				m:     3,
				nums2: []int{2,5,6},
				n:     3,
			},
			want: []int{1,2,2,3,5,6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Solve(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			if !reflect.DeepEqual(tt.args.nums1, tt.want) {
				t.Errorf("Solve() = %v, want %v", tt.args.nums1, tt.want)
			}
		})
	}
}