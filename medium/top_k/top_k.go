package top_k

func TopKFrequent(nums []int, k int) []int {
	numFrequents := map[int]int{}
	frequentNums := map[int][]int{}

	for _, num := range nums {
		numFrequents[num]++
	}

	for num, freq := range numFrequents {
		frequentNums[freq] = append(frequentNums[freq], num)
	}

	topK := []int{}
	for i := len(nums); i > 0; i-- {
		for _, num := range frequentNums[i] {
			if len(topK) == k {
				return topK
			}

			topK = append(topK, num)
		}
	}

	return topK
}
