package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want MinStack
	}{
		{
			name: "normal",
			want: MinStack{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinStack_GetMin(t *testing.T) {
	type fields struct {
		Values []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "[-1,0,1]_return_-1",
			fields: fields{
				Values: []int{-1, 0, 1},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MinStack{
				Values: tt.fields.Values,
			}
			if got := this.GetMin(); got != tt.want {
				t.Errorf("GetMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinStack_Pop(t *testing.T) {
	type fields struct {
		Values []int
	}
	tests := []struct {
		name   string
		fields fields
		want   *MinStack
	}{
		{
			name: "[-1,0,1]_return_[-1,0]",
			fields: fields{
				Values: []int{-1, 0, 1},
			},
			want: &MinStack{
				Values: []int{-1, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MinStack{
				Values: tt.fields.Values,
			}
			this.Pop()
			if !reflect.DeepEqual(this, tt.want) {
				t.Errorf("Pop() = %v, want %v", this, tt.want)
			}
		})
	}
}

func TestMinStack_Push(t *testing.T) {
	type fields struct {
		Values []int
	}
	type args struct {
		x int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *MinStack
	}{
		{
			name: "[-1,0,1]_2_return_[-1,0,1,2]",
			fields: fields{
				Values: []int{-1, 0, 1},
			},
			args: args{
				x: 2,
			},
			want: &MinStack{
				Values: []int{-1, 0, 1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MinStack{
				Values: tt.fields.Values,
			}
			this.Push(tt.args.x)
			if !reflect.DeepEqual(this, tt.want) {
				t.Errorf("Push(%v) = %v, want %v", tt.args.x, this, tt.want)
			}
		})
	}
}

func TestMinStack_Top(t *testing.T) {
	type fields struct {
		Values []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "[-1,0,1]_return_1",
			fields: fields{
				Values: []int{-1, 0, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MinStack{
				Values: tt.fields.Values,
			}
			if got := this.Top(); got != tt.want {
				t.Errorf("Top() = %v, want %v", got, tt.want)
			}
		})
	}
}
