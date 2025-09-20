package functions

import (
	"erew/internal/memory"
)

func SetXToArr(u *memory.Unique, workerId int, xMem *memory.Erew[int], array []*memory.Erew[int]) {
	var x int

	if workerId == 0 {
		x = xMem.Read()
		array[0].Write(x)
	} else {
		u.Skip(2)
	}

	/*
		 Steps from 0 to ceil(log(n)).
		 Invariants on step k:
			1. We are setting segment [2^k, 2^(k + 1)) with length 2^k.
		    2. First 2^k workers writing:
		       worker with id = i writing element i + 2^k
		    3. Next 2^k workers on segment [2^k, 2^(k + 1)) reading array:
		       worker with id = i reading element i - 2^k
		    4. The others workers just skipping.
	*/
	for i := 1; i < len(array); i *= 2 {
		if workerId < i && workerId+i < len(array) {
			array[workerId+i].Write(x)
		} else if workerId >= i && workerId < 2*i {
			x = array[workerId-i].Read()
		} else {
			u.Skip(1)
		}
	}
}
