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
			name: "3 return 0",
			args: args{
				n: 3,
			},
			want: 0,
		},
		{
			name: "5 return 1",
			args: args{
				n: 5,
			},
			want: 1,
		},
		{
			name: "15 return 3",
			args: args{
				n: 15,
			},
			want: 3,
		},
		{
			name: "777 return 193",
			args: args{
				n: 777,
			},
			want: 193,
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
