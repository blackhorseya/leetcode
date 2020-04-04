package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1_return_1",
			args: args{
				n: 1,
			},
			want: "1",
		},
		{
			name: "4_return_1211",
			args: args{
				n: 4,
			},
			want: "1211",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.n); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}