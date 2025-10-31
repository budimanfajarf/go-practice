package main

import (
	"fmt"

	"github.com/budimanfajarf/go-practice/easy"
)

func main() {
	fmt.Println("--- EASY ---")
	fmt.Println("")

	fmt.Println("Number of Steps to Reduce a Number to Zero")
	fmt.Println("NumberOfSteps(14) =", easy.NumberOfSteps(14))   // Output: 6
	fmt.Println("NumberOfSteps(8) =", easy.NumberOfSteps(8))     // Output: 4
	fmt.Println("NumberOfSteps(123) =", easy.NumberOfSteps(123)) // Output: 12
	fmt.Println("")

	fmt.Println("Running Sum of 1d Array")
	fmt.Println("RunningSum([1,2,3,4]) =", easy.RunningSum([]int{1, 2, 3, 4}))        // Output: [1,3,6,10]
	fmt.Println("RunningSum([1,1,1,1,1]) =", easy.RunningSum([]int{1, 1, 1, 1, 1}))   // Output: [1,2,3,4,5]
	fmt.Println("RunningSum([3,1,2,10,1]) =", easy.RunningSum([]int{3, 1, 2, 10, 1})) // Output: [3,4,6,16,17]
	fmt.Println("")

	fmt.Println("Richest Customer Wealth")
	fmt.Println("MaximumWealth([[1,2,3],[3,2,1]]) =", easy.MaximumWealth([][]int{{1, 2, 3}, {3, 2, 1}}))                    // Output: 6
	fmt.Println("MaximumWealth([[1,5],[7,3],[3,5]]) =", easy.MaximumWealth([][]int{{1, 5}, {7, 3}, {3, 5}}))                // Output: 10
	fmt.Println("MaximumWealth([[2,8,7],[7,1,3],[1,9,5]]) =", easy.MaximumWealth([][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}})) // Output: 17
	fmt.Println("")

	fmt.Println("Fizz Buzz")
	fmt.Println("FizzBuzz(3) =", easy.FizzBuzz(3))   // Output: ["1","2","Fizz"]
	fmt.Println("FizzBuzz(5) =", easy.FizzBuzz(5))   // Output: ["1","2","Fizz","4","Buzz"]
	fmt.Println("FizzBuzz(15) =", easy.FizzBuzz(15)) // Output: ["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]
	fmt.Println("")

	fmt.Println("Middle of the Linked List")
	fmt.Println("MiddleNode([1,2,3,4,5]) =", easy.MiddleNode(easy.ListNodeHead1).Val)   // Output: 3
	fmt.Println("MiddleNode([1,2,3,4,5,6]) =", easy.MiddleNode(easy.ListNodeHead2).Val) // Output: 4
	fmt.Println("")

	fmt.Println("--- MEDIUM ---")
	fmt.Println("")

	fmt.Println("TODO")
	fmt.Println("")

	fmt.Println("--- HARD ---")
	fmt.Println("")

	fmt.Println("TODO")
	fmt.Println("")
}
