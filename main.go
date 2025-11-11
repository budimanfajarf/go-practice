package main

import (
	"fmt"

	"github.com/budimanfajarf/go-practice/easy/fizz_buzz"
	"github.com/budimanfajarf/go-practice/easy/largest_good_int"
	"github.com/budimanfajarf/go-practice/easy/maximum_wealth"
	"github.com/budimanfajarf/go-practice/easy/middle_node"
	"github.com/budimanfajarf/go-practice/easy/number_of_steps"
	"github.com/budimanfajarf/go-practice/easy/ransom_note"
	"github.com/budimanfajarf/go-practice/easy/running_sum"
	"github.com/budimanfajarf/go-practice/hard/word_ladder"
	"github.com/budimanfajarf/go-practice/medium/max_subarray"
	"github.com/budimanfajarf/go-practice/medium/top_k"
)

func main() {
	fmt.Println("--- EASY ---")
	fmt.Println("")

	fmt.Println("Running Sum of 1d Array")
	fmt.Println("RunningSum([1,2,3,4]) =", running_sum.RunningSum([]int{1, 2, 3, 4}))        // Output: [1,3,6,10]
	fmt.Println("RunningSum([1,1,1,1,1]) =", running_sum.RunningSum([]int{1, 1, 1, 1, 1}))   // Output: [1,2,3,4,5]
	fmt.Println("RunningSum([3,1,2,10,1]) =", running_sum.RunningSum([]int{3, 1, 2, 10, 1})) // Output: [3,4,6,16,17]
	fmt.Println("")

	fmt.Println("Richest Customer Wealth")
	fmt.Println("MaximumWealth([[1,2,3],[3,2,1]]) =", maximum_wealth.MaximumWealth([][]int{{1, 2, 3}, {3, 2, 1}}))                    // Output: 6
	fmt.Println("MaximumWealth([[1,5],[7,3],[3,5]]) =", maximum_wealth.MaximumWealth([][]int{{1, 5}, {7, 3}, {3, 5}}))                // Output: 10
	fmt.Println("MaximumWealth([[2,8,7],[7,1,3],[1,9,5]]) =", maximum_wealth.MaximumWealth([][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}})) // Output: 17
	fmt.Println("")

	fmt.Println("Fizz Buzz")
	fmt.Println("FizzBuzz(3) =", fizz_buzz.FizzBuzz(3))   // Output: ["1","2","Fizz"]
	fmt.Println("FizzBuzz(5) =", fizz_buzz.FizzBuzz(5))   // Output: ["1","2","Fizz","4","Buzz"]
	fmt.Println("FizzBuzz(15) =", fizz_buzz.FizzBuzz(15)) // Output: ["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]
	fmt.Println("")

	fmt.Println("Number of Steps to Reduce a Number to Zero")
	fmt.Println("NumberOfSteps(14) =", number_of_steps.NumberOfSteps(14))   // Output: 6
	fmt.Println("NumberOfSteps(8) =", number_of_steps.NumberOfSteps(8))     // Output: 4
	fmt.Println("NumberOfSteps(123) =", number_of_steps.NumberOfSteps(123)) // Output: 12
	fmt.Println("")

	fmt.Println("Middle of the Linked List")
	fmt.Println("MiddleNode([1,2,3,4,5]) =", middle_node.MiddleNode(middle_node.Head1).Val)   // Output: 3
	fmt.Println("MiddleNode([1,2,3,4,5,6]) =", middle_node.MiddleNode(middle_node.Head2).Val) // Output: 4
	fmt.Println("")

	fmt.Println("Ransom Note")
	fmt.Println("CanConstruct('a', 'b') =", ransom_note.CanConstruct("a", "b"))         // Output: false
	fmt.Println("CanConstruct('aa', 'ab') =", ransom_note.CanConstruct("aa", "ab"))     // Output: false
	fmt.Println("CanConstruct('aa', 'aab') =", ransom_note.CanConstruct("aa", "aab"))   // Output: true
	fmt.Println("CanConstruct('aab', 'baa') =", ransom_note.CanConstruct("aab", "baa")) // Output: true
	fmt.Println("")

	fmt.Println("Largest 3-Same-Digit Number in String")
	fmt.Println("LargestGoodInteger('6777133339') =", largest_good_int.LargestGoodInteger("6777133339")) // Output: "777"
	fmt.Println("LargestGoodInteger('2300019') =", largest_good_int.LargestGoodInteger("2300019"))       // Output: "000"
	fmt.Println("LargestGoodInteger('42352338') =", largest_good_int.LargestGoodInteger("42352338"))     // Output: ""
	fmt.Println("")

	fmt.Println("--- MEDIUM ---")
	fmt.Println("")

	fmt.Println("Maximum Product Subarray")
	fmt.Println("MaxProduct([2,3,-2,4]) =", max_subarray.MaxProduct([]int{2, 3, -2, 4})) // Output: 6
	fmt.Println("MaxProduct([-2,0,-1]) =", max_subarray.MaxProduct([]int{-2, 0, -1}))    // Output: 0
	fmt.Println("")

	fmt.Println("Top K Frequent Elements")
	fmt.Println("TopKFrequent([1,1,1,2,2,3], 2) =", top_k.TopKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))                     // Output: [1,2]
	fmt.Println("TopKFrequent([1], 1) =", top_k.TopKFrequent([]int{1}, 1))                                              // Output: [1]
	fmt.Println("TopKFrequent([1,2,1,2,1,2,3,1,3,2], 2) =", top_k.TopKFrequent([]int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 2)) // Output: [1,2]
	fmt.Println("TopKFrequent([3,2,1,2], 1) =", top_k.TopKFrequent([]int{3, 2, 1, 2}, 1))                               // Output: [2]
	fmt.Println("")

	fmt.Println("--- HARD ---")
	fmt.Println("")

	fmt.Println("Word Ladder")
	fmt.Println("LadderLength('hit', 'cog', ['hot','dot','dog','lot','log','cog']) =", word_ladder.LadderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})) // Output: 5
	fmt.Println("LadderLength('hit', 'cog', ['hot','dot','dog','lot','log']) =", word_ladder.LadderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))              // Output: 0
	fmt.Println("")
}
