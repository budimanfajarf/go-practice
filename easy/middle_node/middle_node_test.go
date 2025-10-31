package middle_node

import (
	"reflect"
	"testing"
)

func TestMiddleNode(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Case 1",
			args: args{
				head: Head1,
			},
			want: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
		{
			name: "Case 2",
			args: args{
				head: Head2,
			},
			want: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MiddleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MiddleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
