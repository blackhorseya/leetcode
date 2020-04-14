package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		target string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "leet_return_DDR!UURRR!!DDD!",
			args: args{
				target: "leet",
			},
			want: "DDR!UURRR!!DDD!",
		},
		{
			name: "code_return_RR!DDRR!UUL!R!",
			args: args{
				target: "code",
			},
			want: "RR!DDRR!UUL!R!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.target); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}