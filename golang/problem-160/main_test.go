package main

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "[4,1,8,4,5] [5,0,1,8,4,5] return [8,4,5]",
			args: args{
				headA: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 8,
							Next: &ListNode{
								Val:  4,
								Next: &ListNode{5, nil},
							},
						},
					},
				},
				headB: &ListNode{5, &ListNode{0, &ListNode{1, &ListNode{8, &ListNode{4, &ListNode{5, nil}}}}}},
			},
			want: &ListNode{
				Val: 8,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
		{
			name: "[0,9,1,2,4] [3,2,4] return [2,4]",
			args: args{
				&ListNode{0, &ListNode{9, &ListNode{1, &ListNode{2, &ListNode{4, nil}}}}},
				&ListNode{3, &ListNode{2, &ListNode{4, nil}}},
			},
			want: &ListNode{2, &ListNode{4, nil}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
