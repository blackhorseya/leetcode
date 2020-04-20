package main

import "testing"

func TestSolve(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "11 return 3",
			args: args{
				num: 11,
			},
			want: 3,
		},
		{
			name: "128 return 1",
			args: args{
				num: 128,
			},
			want: 1,
		},
		{
			name: "4294967293 return 31",
			args: args{
				num: 4294967293,
			},
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.num); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}