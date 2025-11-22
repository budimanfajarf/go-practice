package most_frequent

func MostFrequent(nums []int, key int) int {
	target := 0
	frequent := 0
	targetFrequents := map[int]int{}

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == key {
			candidate := nums[i+1]
			targetFrequents[candidate]++

			if targetFrequents[candidate] > frequent {
				target = candidate
				frequent = targetFrequents[candidate]
			}
		}
	}

	return target
}
