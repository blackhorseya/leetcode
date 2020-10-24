package main

import "testing"

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
			name: "0 then empty",
			args: args{0},
			want: "",
		},
		{
			name: "1 then A",
			args: args{1},
			want: "A",
		},
		{
			name: "26 then Z",
			args: args{26},
			want: "Z",
		},
		{
			name: "27 then AA",
			args: args{27},
			want: "AA",
		},
		{
			name: "52 then AZ",
			args: args{52},
			want: "AZ",
		},
		{
			name: "701 then ZY",
			args: args{701},
			want: "ZY",
		},
		{
			name: "703 then AAA",
			args: args{703},
			want: "AAA",
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