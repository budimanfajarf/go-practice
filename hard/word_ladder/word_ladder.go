package word_ladder

import "slices"

func LadderLength(beginWord string, endWord string, wordList []string) int {
	if !slices.Contains(wordList, endWord) {
		return 0
	}

	wordList = append(wordList, beginWord)
	wordLength := len(beginWord)
	neighborWords := map[string][]string{}

	for _, word := range wordList {
		for i := range wordLength {
			pattern := word[:i] + "*" + word[i+1:]
			// e.g hot
			// i = 0, pattern = "" + "*" + "ot"
			// i = 1, pattern = "h" + "*" + "t"
			// i = 2, pattern = "ho" + "*" + ""
			neighborWords[pattern] = append(neighborWords[pattern], word)
		}
	}

	queues := []string{}
	isVisited := map[string]bool{}
	queues = append(queues, beginWord)
	isVisited[beginWord] = true
	steps := 1

	for len(queues) > 0 {
		newQueues := []string{}

		for _, word := range queues {
			if word == endWord {
				return steps
			}

			for i := range wordLength {
				pattern := word[:i] + "*" + word[i+1:]

				for _, neighborWord := range neighborWords[pattern] {
					if !isVisited[neighborWord] {
						newQueues = append(newQueues, neighborWord)
						isVisited[neighborWord] = true
					}
				}
			}
		}

		queues = newQueues
		steps++
	}

	return 0
}
