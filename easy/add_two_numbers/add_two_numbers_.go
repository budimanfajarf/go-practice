package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

var L1 = &ListNode{
	Val: 2,
	Next: &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 3,
		},
	},
}

var L2 = &ListNode{
	Val: 5,
	Next: &ListNode{
		Val: 6,
		Next: &ListNode{
			Val: 4,
		},
	},
}
