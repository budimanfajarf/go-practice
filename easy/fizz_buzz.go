package easy

import "strconv"

func FizzBuzz(n int) []string {
	answers := []string{}

	for i := 1; i < n+1; i++ {
		answer := ""
		isDividedBy3 := i%3 == 0
		isDividedBy5 := i%5 == 0

		if isDividedBy3 {
			answer += "Fizz"
		}

		if isDividedBy5 {
			answer += "Buzz"
		}

		if !isDividedBy3 && !isDividedBy5 {
			answer = strconv.Itoa(i)
		}

		answers = append(answers, answer)
	}

	return answers
}
