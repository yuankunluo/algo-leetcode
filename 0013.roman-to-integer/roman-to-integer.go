package problem0013

func romanToInt(s string) int {
	m := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	us := []rune(s)
	var p rune
	ans := 0

	for _, c := range us {
		ans += m[c]
		if p != 0 && m[c] > m[p] {
			ans -= 2 * m[p]
		}
		p = c
	}
	return ans
}
