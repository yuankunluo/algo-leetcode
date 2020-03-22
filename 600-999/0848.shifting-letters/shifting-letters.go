package main

func shiftingLetters(S string, shifts []int) string {
	c := 0
	s := []rune(S)
	for i := len(shifts) - 1; i >= 0; i-- {
		c += (shifts[i] % 26)
		s[i] = (s[i]-'a'+rune(c))%26 + 'a'
	}
	return string(s)
}
