package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "11_1_return_100",
			args: args{
				a: "11",
				b: "1",
			},
			want: "100",
		},
		{
			name: "1010_1011_return_10101",
			args: args{
				a: "1010",
				b: "1011",
			},
			want: "10101",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}