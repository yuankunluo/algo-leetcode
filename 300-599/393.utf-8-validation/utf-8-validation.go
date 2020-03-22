func validUtf8(data []int) bool {
	left := 0
	for _, d := range data {
		if left == 0 {
			if (d >> 3) == 0b11110 {
				left = 3
			} else if (d >> 4) == 0b1110 {
				left = 2
			} else if (d >> 5) == 0b110 {
				left = 1
			} else if (d >> 7) == 0b0 {
				left = 0
			} else {
				return false
			}
		} else {
			if (d >> 6) != 0b10 {
				return false
			}
			left--
		}
	}
	return left == 0
}