func ladderLength(beginWord string, endWord string, wordList []string) int {
	// Create a word map to denote word existence.
	wordMap := make(map[string]bool)

	for _, word := range wordList {
		wordMap[word] = true
	}

	// If there is no endWord then return.
	if !wordMap[endWord] {
		return 0
	}

	beginQueue := []string{beginWord}
	endQueue := []string{endWord}

	isVisited := make(map[string]bool)
	length := 1

	for len(beginQueue) != 0 && len(endQueue) != 0 {
		if len(beginQueue) > len(endQueue) {
			// Switch. We only iterate through shorter queue.
			beginQueue, endQueue = endQueue, beginQueue
		}
		temp := make([]string, 0)
		for _, word := range beginQueue {
			wordChars := []rune(word)

			for i := 0; i< len(wordChars); i++ {
				for c := 'a'; c <= 'z'; c++ {
					old := wordChars[i]
					wordChars[i] = c
					target := string(wordChars)
					if isContain(target, endQueue){
						return length + 1
					}

					if wordMap[target] && !isVisited[target] {
						temp = append(temp, target)
						isVisited[target] = true
					}
					wordChars[i] = old
				}
			}
		}
		beginQueue = temp
		length++
	}
	return 0
}

func isContain(target string, queue []string) bool {
	for _, word := range queue {
		if target == word {
			return true
		}
	}
	return false
}