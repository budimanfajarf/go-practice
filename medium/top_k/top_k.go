package top_k

func TopKFrequent(nums []int, k int) []int {
	numFrequents := map[int]int{}
	frequentNums := map[int][]int{}

	for i := 0; i < len(nums)+1; i++ {
		frequentNums[i+1] = []int{}
	}

	for _, num := range nums {
		numFrequents[num]++
	}

	for num, freq := range numFrequents {
		frequentNums[freq] = append(frequentNums[freq], num)
	}

	topK := []int{}
	for i := len(frequentNums); i > 0; i-- {
		for _, num := range frequentNums[i] {
			if len(topK) == k {
				return topK
			}

			topK = append(topK, num)
		}
	}

	return topK
}
