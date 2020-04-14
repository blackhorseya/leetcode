package main

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		target string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "leet_return_DDR!UURRR!!DDD!",
			args: args{
				target: "leet",
			},
			want: "DDR!UURRR!!DDD!",
		},
		{
			name: "code_return_RR!DDRR!UUL!R!",
			args: args{
				target: "code",
			},
			want: "RR!DDRR!UUL!R!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.target); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPointByLetter(t *testing.T) {
	type args struct {
		letter int32
	}
	tests := []struct {
		name string
		args args
		want *Point
	}{
		{
			name: "l_return_{1,2}",
			args: args{
				letter: 'l',
			},
			want: &Point{
				X: 1,
				Y: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPointByLetter(tt.args.letter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPointByLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_WalkTo(t *testing.T) {
	type fields struct {
		X int32
		Y int32
	}
	type args struct {
		goal *Point
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "{0,0}_{1,2}_return_DDR!",
			fields: fields{
				X: 0,
				Y: 0,
			},
			args:   args{
				goal: &Point{
					X: 1,
					Y: 2,
				},
			},
			want:   "DDR!",
		},
		{
			name:   "{1,2}_{4,0}_return_UURRR!",
			fields: fields{
				X: 1,
				Y: 2,
			},
			args:   args{
				goal: &Point{
					X: 4,
					Y: 0,
				},
			},
			want:   "UURRR!",
		},
		{
			name:   "{4,0}_{4,0}_return_!",
			fields: fields{
				X: 4,
				Y: 0,
			},
			args:   args{
				goal: &Point{
					X: 4,
					Y: 0,
				},
			},
			want:   "!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Point{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := p.WalkTo(tt.args.goal); got != tt.want {
				t.Errorf("WalkTo() = %v, want %v", got, tt.want)
			}
		})
	}
}