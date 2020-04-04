
// Leetcode: Hard
// Tags:  HashTable, TwoPointers, String, Sliding Window
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	// Use map to denote characters.
	wordDict := make(map[rune]int)
	for _, word := range s {
		wordDict[word]++
	}
	// TwoPointers, one is slow, other is fast.
	slow := 0
	minLen := len(s)
	matchCount := 0
	index := 0
	
	for fast := 0; fast < len(s); fast++ {
		word := rune(string[fast])
		count, ok := wordDict[word]
		if !ok {
			continue
		}
		wordDict[word]--
		if wordDict[word] == 1 {
			// 1 -> 0
			matchCount++
		}

		for matchCount == len(wordDict) {
			// find a valid substring.
			if fast - slow + 1 < minLen {
				minLen = fast - slow + 1
				index = slow
			}
			leftmost := rune(string[slow])
			leftmostCount, ok := wordDict[leftmost]
			if !ok {
				slow++
				continue
			} else {
				wordDict[leftmost]++
				if leftmostCount
			}
		}
	}

}