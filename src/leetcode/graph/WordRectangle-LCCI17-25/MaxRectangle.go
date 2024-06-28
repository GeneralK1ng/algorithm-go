package WordRectangle_LCCI17_25

type Trie struct {
	child [26]*Trie
	isEnd bool
}

func (this *Trie) inputTrie(word string) {
	node := this
	for i := 0; i < len(word); i++ {
		u := int(word[i] - 'a')
		if node.child[u] == nil {
			node.child[u] = &Trie{[26]*Trie{}, false}
		}
		if i == len(word)-1 {
			node.child[u].isEnd = true
			return
		}
		node = node.child[u]
	}

}

func (this *Trie) findPrefix(word string) bool {
	child := this.child
	for i := 0; i < len(word); i++ {
		u := word[i] - 'a'
		if child[u] == nil && child[u].isEnd {
			return true
		}
		child = child[u].child
	}
	return false

}

func maxRectangle(words []string) []string {
	trie := &Trie{[26]*Trie{}, false}
	for i := 0; i < len(words); i++ {
		trie.inputTrie(words[i])
	}
	maxLength := 0
	m := map[int][]string{}
	for i := 0; i < len(words); i++ {
		length := len(words[i])
		maxLength = getMax(length, maxLength)
		m[length] = append(m[length], words[i])
	}

	maxArea := 0
	ans := []string{}

	dfs := func(path []string, wordList []string, length int) {}
	dfs = func(path []string, wordList []string, length int) {
		if length*maxLength <= maxArea {
			return
		}
		if len(path) > maxLength {
			return
		}

		for i := 0; i < len(wordList); i++ {
			path = append(path, wordList[i])
			valid, isLeaf := isValid(path, trie)
			if valid {
				area := len(path) * len(path[0])
				if isLeaf && area > maxArea {
					maxArea = area
					ans = make([]string, len(path))
					copy(ans, path)

				}
				dfs(path, wordList, length)
			}
			path = path[:len(path)-1]
		}

	}
	for key, value := range m {
		dfs([]string{}, value, key)
	}

	return ans
}

func isValid(path []string, trie *Trie) (bool, bool) {
	length := len(path[0])
	columnSize := len(path)
	isLeaf := true
	for j := 0; j < length; j++ {
		node := trie
		for i := 0; i < columnSize; i++ {
			u := path[i][j] - 'a'
			child := node.child[u]
			if child == nil {
				return false, false
			}
			node = child
		}
		if !node.isEnd {
			isLeaf = false
		}
	}
	return true, isLeaf
}

func getMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
