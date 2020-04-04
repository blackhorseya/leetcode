package main

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		first  *ListNode
		second *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "two nil ListNode",
			args: args{
				first:  nil,
				second: nil,
			},
			want: nil,
		},
		{
			name: "any nil ListNode",
			args: args{
				first: nil,
				second: &ListNode{
					Val:  0,
					Next: nil,
				},
			},
			want: nil,
		},
		{
			name: "input: (2 -> 4 -> 3) + (5 -> 6 -> 4), output: (7 -> 0 -> 8)",
			args: args{
				first: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
				second: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: &ListNode{
				Val: 8,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  7,
						Next: nil,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.first, tt.args.second); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListNode_ToRevInt(t *testing.T) {
	type fields struct {
		Val  int
		Next *ListNode
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "2->4->3",
			fields: fields{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
			want: 342,
		},
		{
			name: "5->6->4",
			fields: fields{
				Val: 5,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			want: 465,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &ListNode{
				Val:  tt.fields.Val,
				Next: tt.fields.Next,
			}
			if got := l.ToRevInt(); got != tt.want {
				t.Errorf("ToRevInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntToListNode(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "0",
			args: args{
				number: 0,
			},
			want: &ListNode{
				Val:  0,
				Next: nil,
			},
		},
		{
			name: "807",
			args: args{
				number: 807,
			},
			want: &ListNode{
				Val:  8,
				Next: &ListNode{
					Val:  0,
					Next: &ListNode{
						Val:  7,
						Next: nil,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToListNode(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntToListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}