package problem0139


func wordBreak(s string, wordDict []string) bool {
	// Covert the wordDict to map
	dict := make(map[string]bool)
	for _, s := range(wordDict) {
		dict[s] = true
	}

	// Create a memory map
	mem := make(map[string]bool)
	// Call recursively
	return wordBreakRec(s, mem, dict)
}

func wordBreakRec(s string, mem, dict map[string]bool) bool {
	// if in memory, direct return
	if mem[s] {
		return mem[s]
	}
	// if the word is in dictionary, memory it and return
	if dict[s]{
		msm[s] = true
		return mem[s]
	}
	// else try every word break 
	for i:=0;i<len(s); i++ {
		left := s[0:i]
		right := s[i:]
		if (dict[right] && wordBreakRec(left, mem, dict)){
			mem[s] = true 
			return mem[s]
		}
	}
	// No solution
	mem[s] = false
	return false
}

