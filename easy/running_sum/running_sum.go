package running_sum

func RunningSum(nums []int) []int {
	sum := 0
	for i := range nums {
		sum = sum + nums[i]
		nums[i] = sum
	}
	return nums
}
