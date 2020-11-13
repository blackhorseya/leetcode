package main

import "testing"

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
			name: "empty then 0",
			args: args{""},
			want: 0,
		},
		{
			name: "space then 0",
			args: args{"    "},
			want: 0,
		},
		{
			name: "[words and 987] then 0",
			args: args{"words and 987"},
			want: 0,
		},
		{
			name: "[   -42] then -42",
			args: args{"   -42"},
			want: -42,
		},
		{
			name: "[4193 with words] then 4193",
			args: args{"4193 with words"},
			want: 4193,
		},
		{
			name: "[-91283472332] then -2147483648",
			args: args{"-91283472332"},
			want: -2147483648,
		},
		{
			name: "[42] then 42",
			args: args{"42"},
			want: 42,
		},
		{
			name: "[-5-] then -5",
			args: args{"-5-"},
			want: -5,
		},
		{
			name: "[+-12] then 0",
			args: args{"+-12"},
			want: 0,
		},
		{
			name: "[+1] then 1",
			args: args{"+1"},
			want: 1,
		},
		{
			name: "[9223372036854775808] then 2147483648",
			args: args{"9223372036854775808"},
			want: 2147483647,
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
