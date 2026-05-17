package top_k

func TopKFrequent(nums []int, k int) []int {
	numFrequents := map[int]int{}
	frequentNums := map[int][]int{}

	// Count the frequency of each number.
	for _, num := range nums {
		numFrequents[num]++
	}

	// Group numbers by their frequency.
	for num, freq := range numFrequents {
		frequentNums[freq] = append(frequentNums[freq], num)
	}

	// Iterate through the possible top frequencies from the maximum of "nums" in descending order,
	// and collect the top frequent numbers until we have "k" of them.
	topK := []int{}
	for i := len(nums); i > 0; i-- {
		numsInFrequent, exist := frequentNums[i]
		if !exist {
			continue
		}

		for _, num := range numsInFrequent {
			if len(topK) == k {
				return topK
			}

			topK = append(topK, num)
		}
	}

	return topK
}
