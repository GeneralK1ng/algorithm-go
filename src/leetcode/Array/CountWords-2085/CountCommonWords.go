package CountWords_2085

func countWords(words1 []string, words2 []string) int {
	wordCountMap := make(map[string]int, len(words1)*2)

	for _, word := range words1 {
		wordCountMap[word]++
	}

	result := 0

	for _, word := range words2 {
		count, exists := wordCountMap[word]

		if !exists {
			continue
		}
		if count == -1 {
			wordCountMap[word] = 0
			result--
		}
		if count == 1 {
			wordCountMap[word] = -1
			result++
		}
	}

	return result
}
