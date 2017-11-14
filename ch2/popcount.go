package ch2

var pc [256]byte

func init() {
	for i, _ := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCountByTable xxx
func PopCountByTable(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCountByLoop xxx
func PopCountByLoop(x uint64) int {
	var bit uint64 = 0x1
	count := 0
	for i := 0; i < 64; i++ {
		if (x&bit) > 0 {
			count++
		}
		bit <<= 1
	}
	return count
}

// PopCountByBit xxx
func PopCountByBit(x uint64) int {
	count := 0
	for x > 0 {
		x &= (x-1)
		count++
	}
	return count
}
