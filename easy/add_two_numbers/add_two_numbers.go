package add_two_numbers

import (
	"fmt"
	"strconv"
)

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1ConcatenatedNumber := concateNumbers(l1, "")
	l2ConcatenatedNumber := concateNumbers(l2, "")

	l1Number, _ := strconv.Atoi(l1ConcatenatedNumber)
	l2Number, _ := strconv.Atoi(l2ConcatenatedNumber)

	sum := l1Number + l2Number
	sumStr := strconv.Itoa(sum)

	return convertStrNumberToListNode(sumStr, 0, nil)
}

func concateNumbers(l *ListNode, concatenatedNumber string) string {
	if l.Next == nil {
		return fmt.Sprintf("%d%s", l.Val, concatenatedNumber)
	}

	concatenatedNumber = fmt.Sprintf("%d%s", l.Val, concatenatedNumber)
	return concateNumbers(l.Next, concatenatedNumber)
}

func convertStrNumberToListNode(strNumber string, count int, l *ListNode) *ListNode {
	if count == len(strNumber) {
		return l
	}

	value, _ := strconv.Atoi(string(strNumber[count]))
	l = &ListNode{
		Val:  value,
		Next: l,
	}

	count++
	return convertStrNumberToListNode(strNumber, count, l)
}
