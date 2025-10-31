package main

import (
	"fmt"

	"github.com/budimanfajarf/go-practice/easy"
)

func main() {
	// Easy
	fmt.Println("Number of Steps to Reduce a Number to Zero")
	fmt.Println("NumberOfSteps(14) =", easy.NumberOfSteps(14))   // Output: 6
	fmt.Println("NumberOfSteps(8) =", easy.NumberOfSteps(8))     // Output: 4
	fmt.Println("NumberOfSteps(123) =", easy.NumberOfSteps(123)) // Output: 12
}
