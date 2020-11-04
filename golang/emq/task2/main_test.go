package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "CBACD then C",
			args: args{"CBACD"},
			want: "C",
		},
		{
			name: "CABABD then empty",
			args: args{"CABABD"},
			want: "",
		},
		{
			name: "ACBDACBD then ACBDACBD",
			args: args{"ACBDACBD"},
			want: "ACBDACBD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.S); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
