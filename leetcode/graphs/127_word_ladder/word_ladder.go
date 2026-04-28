package word_ladder

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)

	for _, word := range wordList {
		wordSet[word] = true
	}

	if _, exists := wordSet[endWord]; !exists {
		return 0
	}

	wordQueue := []string{beginWord}

	distance := 1

	for len(wordQueue) > 0 {
		levelSize := len(wordQueue)

		for i := 0; i < levelSize; i++ {
			currentWord := wordQueue[0]
			wordQueue = wordQueue[1:]

			if currentWord == endWord {
				return distance
			}

			for j := 0; j < len(currentWord); j++ {
				for c := 'a'; c <= 'z'; c++ {
					if byte(c) == currentWord[j] {
						continue
					}

					newWord := currentWord[:j] + string(c) + currentWord[j+1:]

					if _, exists := wordSet[newWord]; exists {
						wordQueue = append(wordQueue, newWord)
						delete(wordSet, newWord)
					}
				}
			}
		}

		distance++
	}

	return 0
}
