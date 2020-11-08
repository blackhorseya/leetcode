package main

import "testing"

func TestSolve(t *testing.T) {
	type args struct {
		dividend int
		divisor  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0 1 then 0",
			args: args{0, 1},
			want: 0,
		},
		{
			name: "1 1 then 1",
			args: args{1, 1},
			want: 1,
		},
		{
			name: "10 3 then 3",
			args: args{10, 3},
			want: 3,
		},
		{
			name: "7 -3 then -2",
			args: args{7, -3},
			want: -2,
		},
		{
			name: "-2147483648 -1 then 2147483647",
			args: args{-2147483648, -1},
			want: 2147483647,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.dividend, tt.args.divisor); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
