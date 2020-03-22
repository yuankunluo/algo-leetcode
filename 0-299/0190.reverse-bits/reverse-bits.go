package problem0190

func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		// num & 1 ：get the last bit form num.
		// eg: 00101 & 1
		res = (res << 1) | (num & 1)
		num = num >> 1
	}
	return res
}
