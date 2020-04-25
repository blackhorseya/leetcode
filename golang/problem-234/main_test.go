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
		want bool
	}{
		{
			name: "nil return true",
			args: args{},
			want: true,
		},
		{
			name: "1 return true",
			args: args{
				head: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
			want: true,
		},
		{
			name: "1->2 return false",
			args: args{
				head: &ListNode{
					Val:  1,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
			},
			want: false,
		},
		{
			name: "1->2->2->1 return true",
			args: args{
				head: &ListNode{
					Val:  1,
					Next: &ListNode{
						Val:  2,
						Next: &ListNode{
							Val:  2,
							Next: &ListNode{
								Val:  1,
								Next: nil,
							},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "1->2->1 return true",
			args: args{
				head: &ListNode{
					Val:  1,
					Next: &ListNode{
						Val:  2,
						Next: &ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.head); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	type fields struct {
		Val  int
		Next *ListNode
	}
	tests := []struct {
		name   string
		fields fields
		want   *ListNode
	}{
		{
			name:   "nil return nil",
			fields: fields{},
			want:   &ListNode{},
		},
		{
			name:   "1 return 1",
			fields: fields{1, nil},
			want:   &ListNode{1, nil},
		},
		{
			name:   "1->2->3 return 3->2->1",
			fields: fields{
				Val:  1,
				Next: &ListNode{
					Val:  2,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
			want:   &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  2,
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
			n := &ListNode{
				Val:  tt.fields.Val,
				Next: tt.fields.Next,
			}
			if got := reverse(n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}