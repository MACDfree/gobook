package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount count bit 1
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCount25 count bit 1 of x
func PopCount25(x uint64) int {
	count := 0
	for ; x != 0; x = x & (x - 1) {
		count++
	}
	return count
}

func PopCountByts(buf [32]byte) int {
	count := 0
	for _, x := range buf {
		for ; x != 0; x = x & (x - 1) {
			count++
		}
	}

	return count
}
