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
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
		{
			name: "input: (5) + (5), output: (0->1)",
			args: args{
				first: &ListNode{
					Val: 5,
					Next: nil,
				},
				second: &ListNode{
					Val: 5,
					Next: nil,
				},
			},
			want: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 1,
					Next: nil,
				},
			},
		},
		{
			name: "input: (1->8) + (0), output: (1->8)",
			args: args{
				first: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
				second: &ListNode{
					Val: 0,
					Next: nil,
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 8,
					Next: nil,
				},
			},
		},
		{
			name: "input: (9->8) + (1), output: (0->9)",
			args: args{
				first: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
				second: &ListNode{
					Val: 1,
					Next: nil,
				},
			},
			want: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 9,
					Next: nil,
				},
			},
		},
		{
			name: "input: (1) + (9->9), output: (0->0->1)",
			args: args{
				first: &ListNode{
					Val: 1,
					Next: nil,
				},
				second: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val:  9,
						Next: nil,
					},
				},
			},
			want: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  1,
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

func TestToSlice(t *testing.T) {
	type fields struct {
		Val  int
		Next *ListNode
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			name: "8->0->7",
			fields: fields{
				Val: 8,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  7,
						Next: nil,
					},
				},
			},
			want: []int{8, 0, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &ListNode{
				Val:  tt.fields.Val,
				Next: tt.fields.Next,
			}
			if got := ToSlice(l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
