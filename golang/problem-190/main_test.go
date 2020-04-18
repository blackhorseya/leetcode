package main

import (
	"math"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "0 return 0",
			args: args{
				num: 0,
			},
			want: 0,
		},
		{
			name: "max return max",
			args: args{
				num: math.MaxUint32,
			},
			want: math.MaxUint32,
		},
		{
			name: "1 return 2147483648",
			args: args{
				num: 1,
			},
			want: 2147483648,
		},
		{
			name: "43261596 return 964176192",
			args: args{
				num: 43261596,
			},
			want: 964176192,
		},
		{
			name: "4294967293 return 3221225471",
			args: args{
				num: 4294967293,
			},
			want: 3221225471,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.num); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}