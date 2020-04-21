package main

import "testing"

func TestSolve(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1 return true",
			args: args{
				n: 1,
			},
			want: true,
		},
		{
			name: "0 return false",
			args: args{
				n: 0,
			},
			want: false,
		},
		{
			name: "19 return true",
			args: args{
				n: 19,
			},
			want: true,
		},
		{
			name: "18 return false",
			args: args{
				n: 18,
			},
			want: false,
		},
		{
			name: "7 return true",
			args: args{
				n: 7,
			},
			want: true,
		},
		{
			name: "13 return true",
			args: args{
				n: 13,
			},
			want: true,
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