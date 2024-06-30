package RemoveLetterEqualizeFreq_2423

func equalFrequency(word string) bool {
	m := make(map[byte]int)
	for i := 0; i < len(word); i++ {
		m[word[i]]++
	}
	for i := 0; i < len(word); i++ {
		m[word[i]]--
		if m[word[i]] == 0 {
			delete(m, word[i])
		}
		if len(m) == 1 {
			return true
		}
	}
	return false
}
