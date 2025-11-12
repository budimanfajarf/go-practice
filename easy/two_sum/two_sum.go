package two_sum

func TwoSum(nums []int, target int) []int {
	mapNums := map[int]int{}

	for i, num := range nums {
		lookedNum := target - num
		if value, ok := mapNums[lookedNum]; ok {
			return []int{value, i}
		}

		mapNums[num] = i
	}

	return nil
}
