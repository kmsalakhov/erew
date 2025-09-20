package task_2

import (
	"fmt"
	"testing"

	"erew/internal/memory"
)

func TestTask(t *testing.T) {
	t.Run("smoke", func(t *testing.T) {
		m := memory.NewManager(5)

		table := make([][]*memory.Erew[int], 5)
		for i := range 5 {
			table[i] = make([]*memory.Erew[int], 5)
			for j := range 5 {
				table[i][j] = memory.AllocateMemoryWithData(m, j)
			}
		}
		printTable(m, table)

		lambda := func(u *memory.Unique, workerId int, args ...interface{}) {
			GetTurtle(u, workerId, args[0].([][]*memory.Erew[int]))
		}

		m.Run(lambda, table)

		printTable(m, table)
	})
}

func printTable(m *memory.Manager, table [][]*memory.Erew[int]) {
	for i := range 5 {
		slice := memory.GetDataSlice(m, table[i])
		fmt.Println(slice)
	}
	fmt.Println()
}
