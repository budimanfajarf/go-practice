package largest_good_int

import (
	"fmt"
)

func LargestGoodInteger(num string) string {
	// Approach 1: Iterate through the string and check for 3 consecutive same digits.
	// largestGoodInt := ""

	// var prevChar rune
	// count := 0
	// for _, char := range num {
	// 	if char == prevChar {
	// 		count++
	// 	} else {
	// 		count = 1
	// 	}

	// 	if count == 3 {
	// 		goodInt := fmt.Sprintf("%c%c%c", char, char, char)

	// 		if goodInt > largestGoodInt {
	// 			largestGoodInt = goodInt
	// 		}
	// 	}

	// 	prevChar = char
	// }

	// return largestGoodInt

	// Approach 2: Sliding window of size 3
	// result := ""
	// for i := 0; i < len(num)-2; i++ {
	// 	if num[i] == num[i+1] && num[i] == num[i+2] {
	// 		goodInt := num[i : i+3]
	// 		if goodInt > result {
	// 			result = goodInt
	// 		}
	// 	}
	// }
	// return result

	// Approach 3
	var largestGoodInt byte

	for i := 0; i < len(num)-2; i++ {
		if num[i] == num[i+1] && num[i] == num[i+2] && num[i] > largestGoodInt {
			largestGoodInt = num[i]
		}
	}

	if largestGoodInt == 0 {
		return ""
	}

	return fmt.Sprintf("%c%c%c", largestGoodInt, largestGoodInt, largestGoodInt)
}
