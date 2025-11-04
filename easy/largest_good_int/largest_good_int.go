package largest_good_int

import "fmt"

func LargestGoodInteger(num string) string {
	largestGoodInt := ""

	var prevChar rune
	count := 0
	for _, char := range num {
		if char == prevChar {
			count++
		} else {
			count = 1
		}

		if count == 3 {
			goodInt := fmt.Sprintf("%c%c%c", char, char, char)

			if goodInt > largestGoodInt {
				largestGoodInt = goodInt
			}
		}

		prevChar = char
	}

	return largestGoodInt
}
