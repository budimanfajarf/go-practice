package middle_node

func MiddleNode(head *ListNode) *ListNode {
	nodes := []ListNode{}
	node := head

	for node != nil {
		nodes = append(nodes, *node)
		node = node.Next
	}

	return &nodes[len(nodes)/2]
}
