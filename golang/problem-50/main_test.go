package main

import "testing"

func TestSolve(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "2 10 then 1024",
			args: args{2, 10},
			want: 1024,
		},
		{
			name: "2.1 3 then 9.261",
			args: args{2.1, 3},
			want: 9.261,
		},
		{
			name: "2 -2 then 0.25",
			args: args{2, -2},
			want: 0.25,
		},
		{
			name: "8.88023 3 then 700.28148",
			args: args{8.88023, 3},
			want: 700.28148,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
