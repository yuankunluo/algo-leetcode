package main

import (
	"fmt"
)

func isPalindrome(s string) bool {

	su := []rune(s)
	n := len(su)
	// Empty and only one char is palindrome.
	if n <= 1 {
		return true
	}

	i := 0
	j := n - 1

	for i < j {
		ci := su[i]
		cj := su[j]

		if !isAlphabet(ci) && !isNumber(ci) {
			i++
			continue
		}
		if !isAlphabet(cj) && !isNumber(cj) {
			j--
			continue
		}

		if isNumber(ci) && ci != cj {
			return false
		}

		if isAlphabet(ci) && isAlphabet(cj) {
			if charEqual(ci, cj) {
				return true
			}
			return false
		} else {
			return false
		}
		i++
		j--
	}

	return true
}

func isAlphabet(c rune) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z'
}

func isNumber(c rune) bool {
	return c >= '0' && c <= '9'
}

// The distance between (A, a) is 32
func charEqual(a, b rune) bool {
	return a-b == 0 || a-b == 32 || a-b == -32
}

func main() {
	input := "A man, a plan, a canal: Panama"

	fmt.Println(isPalindrome(input))
}
