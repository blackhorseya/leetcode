package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "III_return_3",
			args: args{
				s: "III",
			},
			want: 3,
		},
		{
			name: "IV_return_4",
			args: args{
				s: "IV",
			},
			want: 4,
		},
		{
			name: "IX_return_9",
			args: args{
				s: "IX",
			},
			want: 9,
		},
		{
			name: "LVIII_return_58",
			args: args{
				s: "LVIII",
			},
			want: 58,
		},
		{
			name: "MCMXCIV_return_1994",
			args: args{
				s: "MCMXCIV",
			},
			want: 1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.s); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}