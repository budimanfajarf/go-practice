package max_subarray

import "slices"

func MaxProduct(nums []int) int {
	result := slices.Max(nums)
	maxSoFar := 1
	minSoFar := 1

	for _, num := range nums {
		tmpMax := maxSoFar * num
		tmpMin := minSoFar * num
		maxSoFar = max(tmpMax, tmpMin, num)
		minSoFar = min(tmpMax, tmpMin, num)
		result = max(result, maxSoFar)
	}

	return result
}
