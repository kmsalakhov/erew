package main

import (
	"erew/internal/functions/task_1"
	"fmt"

	"erew/internal/memory"
)

const workers = 10

func main() {
	m := memory.NewManager(workers)

	lambda := func(u *memory.Unique, workerId int, args ...interface{}) {
		task_1.SetXToArr(u, workerId, args[0].(*memory.Erew[int]), args[1].([]*memory.Erew[int]))
	}

	x := memory.AllocateMemoryWithData[int](m, 123)
	array := memory.AllocateMemorySlice[int](m, workers)

	m.Run(lambda, x, array)

	fmt.Println(memory.GetDataSlice(m, array))
}
