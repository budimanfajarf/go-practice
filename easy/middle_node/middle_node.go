package middle_node

func MiddleNode(head *ListNode) *ListNode {
	// Approach 1: Store nodes in a slice and return the middle one
	// nodes := []ListNode{}

	// for head != nil {
	// 	nodes = append(nodes, *head)
	// 	head = head.Next
	// }

	// return &nodes[len(nodes)/2]

	// Approach 2: Use two pointers (middle and end), and use end pointer to end the loop
	middle, end := head, head

	for end != nil && end.Next != nil {
		middle = middle.Next
		end = end.Next.Next
	}

	return middle
}
