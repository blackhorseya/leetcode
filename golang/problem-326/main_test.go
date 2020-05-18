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
			name: "0 return false",
			args: args{0},
			want: false,
		},
		{
			name: "27 return true",
			args: args{27},
			want: true,
		},
		{
			name: "9 return true",
			args: args{9},
			want: true,
		},
		{
			name: "45 return false",
			args: args{45},
			want: false,
		},
		{
			name: "1 return true",
			args: args{1},
			want: true,
		},
		{
			name: "3 return true",
			args: args{3},
			want: true,
		},
		{
			name: "19684 return false",
			args: args{19684},
			want: false,
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
