package examples

import "time"

func fibb(a uint) uint {
	if a == 0 {
		return 0
	}

	if a == 1 {
		return 1
	}

	return fibb(a-1) + fibb(a-2)
}

func after() {
	ch := time.After(1 * time.Second)
	<-ch
}

func allocLots() []int {
	arr := []int{}
	for i := 0; i < 1000000; i++ {
		arr = append(arr, i)
	}
	return arr
}
