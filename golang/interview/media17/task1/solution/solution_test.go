package solution

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "24 then",
			args: args{24},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Solution(tt.args.N)
		})
	}
}
