package main

import (
	"fmt"

	"erew/internal/functions"
	"erew/internal/memory"
)

const workers = 125

func main() {
	m := memory.NewManager(workers)

	// TODO(kasalakhov): think about auto cast (with reflection?)
	lambda := func(u *memory.Unique, workerId int, args ...interface{}) {
		functions.SetXToArr(u, workerId, args[0].(*memory.Erew[int]), args[1].([]*memory.Erew[int]))
	}

	x := memory.AllocateMemoryWithData[int](m, 101)
	array := memory.AllocateMemorySlice[int](m, workers)

	m.Run(lambda, x, array)

	fmt.Println(memory.GetDataSlice(m, array))
}
