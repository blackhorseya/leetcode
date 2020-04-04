package main

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "nil_return_nil",
			args: args{
				head: nil,
			},
			want: nil,
		},
		{
			name: "1->2->3->4_return_2->1->4->3",
			args: args{
				head: &ListNode{
					Val:  1,
					Next: &ListNode{
						Val:  2,
						Next: &ListNode{
							Val:  3,
							Next: &ListNode{
								Val:  4,
								Next: nil,
							},
						},
					},
				},
			},
			want: &ListNode{
				Val:  2,
				Next: &ListNode{
					Val:  1,
					Next: &ListNode{
						Val:  4,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}