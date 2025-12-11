// https://youtu.be/dOonV4byDEg?si=Jd_gCsaqGFcfy1sB
package main

import "fmt"

func maxWindowSumSubarrayOfK(nums []int, k int) int {
	fmt.Print("nums: ", nums, " | ")
	fmt.Print("k: ", k, "\n")

	left := 0
	windowSum := 0
	maxWindowSum := 0

	for i := range k {
		windowSum += nums[i]
	}
	maxWindowSum = windowSum

	fmt.Print("initial window sum: ", windowSum, " | ")
	fmt.Print("initial max window sum: ", maxWindowSum, "\n")

	for right := k; right < len(nums); right++ {
		leftToSubstract := nums[left]
		rightToAdd := nums[right]
		windowSum = windowSum - leftToSubstract + rightToAdd

		fmt.Print("left to substract: ", leftToSubstract, " | ")
		fmt.Print("right to add: ", rightToAdd, " | ")
		fmt.Print("window sum: ", windowSum, "\n")

		left++
		maxWindowSum = max(maxWindowSum, windowSum)
	}

	return maxWindowSum
}

func main() {
	result := maxWindowSumSubarrayOfK([]int{8, 3, -2, 4, 5, -1, 0, 5, 3, 9, 6}, 5)
	fmt.Println("result max window sum:", result)
}
