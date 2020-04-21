package main

import "testing"

func TestSolve(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0 return 0",
			args: args{0},
			want: 0,
		},
		{
			name: "1 return 0",
			args: args{1},
			want: 0,
		},
		{
			name: "2 return 0",
			args: args{2},
			want: 0,
		},
		{
			name: "3 return 1",
			args: args{3},
			want: 1,
		},
		{
			name: "10 return 4",
			args: args{10},
			want: 4,
		},
		{
			name: "100 return 25",
			args: args{100},
			want: 25,
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