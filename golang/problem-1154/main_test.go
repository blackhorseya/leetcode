package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty_return_0",
			args: args{},
			want: 0,
		},
		{
			name: "error_format_return_0",
			args: args{
				date: "2342131547",
			},
			want: 0,
		},
		{
			name: "2019-01-09_return_9",
			args: args{
				date: "2019-01-09",
			},
			want: 9,
		},
		{
			name: "2019-02-10_return_41",
			args: args{
				date: "2019-02-10",
			},
			want: 41,
		},
		{
			name: "2003-03-01_return_60",
			args: args{
				date: "2003-03-01",
			},
			want: 60,
		},
		{
			name: "2004-03-01_return_61",
			args: args{
				date: "2004-03-01",
			},
			want: 61,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.date); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
