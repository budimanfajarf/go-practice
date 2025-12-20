// https://www.perplexity.ai/search/func-addtwonumbers-l1-listnode-cvCwwbnJT4.DriFhVYYCbw#2
package add_two_numbers

import "fmt"

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var prevHead *ListNode
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry

		if l1 != nil {
			fmt.Print("l1.Val:", l1.Val, " | ")

			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			fmt.Print("l2.Val:", l2.Val, " | ")

			sum += l2.Val
			l2 = l2.Next
		}

		fmt.Print("sum:", sum, " | ")

		newNode := &ListNode{
			Val: sum % 10,
		}

		fmt.Print("newNode:", newNode, " | ")

		if head == nil {
			// head is empty, assign the new node(*pointer) to head & prevHead
			head = newNode
			prevHead = newNode
		} else {
			prevHead.Next = newNode // assign the new node(*pointer) to the previous head's Next
			prevHead = newNode      // replace the prevHead with the new node(*pointer)
		}

		carry = sum / 10
		fmt.Print("carry:", carry, "\n")
	}

	return head
}
