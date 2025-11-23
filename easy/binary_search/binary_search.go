package binary_search

func Search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := low + (high-low)/2
		num := nums[mid]

		if target == num {
			return mid
		} else if target > num {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
