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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.s); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDigit(t *testing.T) {
	type args struct {
		c rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "A then false",
			args: args{'A'},
			want: false,
		},
		{
			name: "( then false",
			args: args{'('},
			want: false,
		},
		{
			name: "0 then true",
			args: args{'0'},
			want: true,
		},
		{
			name: "9 then true",
			args: args{'9'},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDigit(tt.args.c); got != tt.want {
				t.Errorf("IsDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSign(t *testing.T) {
	type args struct {
		c rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "A then false",
			args: args{'A'},
			want: false,
		},
		{
			name: "( then false",
			args: args{'('},
			want: false,
		},
		{
			name: "0 then false",
			args: args{'0'},
			want: false,
		},
		{
			name: "9 then false",
			args: args{'9'},
			want: false,
		},
		{
			name: "+ then true",
			args: args{'+'},
			want: true,
		},
		{
			name: "- then true",
			args: args{'-'},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSign(tt.args.c); got != tt.want {
				t.Errorf("IsSign() = %v, want %v", got, tt.want)
			}
		})
	}
}