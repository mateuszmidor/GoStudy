package popcount

var pc [256]byte

// init is started automatically when the package is initialized for run
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the populatin count (number of bits set) for x
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

func popCount2(x int64) int {
	count := 0
	for i := uint(0); i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
}
