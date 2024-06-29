package TwoEditWord_2452

func twoEditWords(queries []string, dictionary []string) []string {
	var res []string
	dictSet := make(map[string]bool)
	for _, word := range dictionary {
		dictSet[word] = true
	}

	for _, query := range queries {
		if canTransform(query, dictSet) {
			res = append(res, query)
		}
	}

	return res
}

func canTransform(word string, dict map[string]bool) bool {
	if dict[word] {
		return true
	}

	n := len(word)
	for target := range dict {
		if len(target) != n {
			continue
		}
		diffCount := 0
		for i := 0; i < n; i++ {
			if word[i] != target[i] {
				diffCount++
				if diffCount > 2 {
					break
				}
			}
		}
		if diffCount <= 2 {
			return true
		}
	}
	return false
}
