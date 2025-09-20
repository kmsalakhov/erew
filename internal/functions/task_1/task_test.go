package task_1

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"erew/internal/memory"
)

const elem = 123456

func TestTask(t *testing.T) {
	t.Run("small workers count", func(t *testing.T) {
		const workers = 10
		expected := slices.Repeat([]int{elem}, workers)

		m := memory.NewManager(workers)

		lambda := func(u *memory.Unique, workerId int, args ...interface{}) {
			SetXToArr(u, workerId, args[0].(*memory.Erew[int]), args[1].([]*memory.Erew[int]))
		}

		x := memory.AllocateMemoryWithData[int](m, elem)
		array := memory.AllocateMemorySlice[int](m, workers)

		m.Run(lambda, x, array)
		actual := memory.GetDataSlice(m, array)
		assert.Equal(t, expected, actual)

		fmt.Println(memory.GetDataSlice(m, array))
	})

	t.Run("big workers count", func(t *testing.T) {
		const workers = 10_000
		expected := slices.Repeat([]int{elem}, workers)

		m := memory.NewManager(workers)

		lambda := func(u *memory.Unique, workerId int, args ...interface{}) {
			SetXToArr(u, workerId, args[0].(*memory.Erew[int]), args[1].([]*memory.Erew[int]))
		}

		x := memory.AllocateMemoryWithData[int](m, elem)
		array := memory.AllocateMemorySlice[int](m, workers)

		m.Run(lambda, x, array)
		actual := memory.GetDataSlice(m, array)
		assert.Equal(t, expected, actual)
	})

	t.Run("one worker", func(t *testing.T) {
		const workers = 1
		expected := slices.Repeat([]int{elem}, workers)

		m := memory.NewManager(workers)

		lambda := func(u *memory.Unique, workerId int, args ...interface{}) {
			SetXToArr(u, workerId, args[0].(*memory.Erew[int]), args[1].([]*memory.Erew[int]))
		}

		x := memory.AllocateMemoryWithData[int](m, elem)
		array := memory.AllocateMemorySlice[int](m, workers)

		m.Run(lambda, x, array)
		actual := memory.GetDataSlice(m, array)
		assert.Equal(t, expected, actual)
	})

	t.Run("two power worker", func(t *testing.T) {
		const workers = 1 << 8
		expected := slices.Repeat([]int{elem}, workers)

		m := memory.NewManager(workers)

		lambda := func(u *memory.Unique, workerId int, args ...interface{}) {
			SetXToArr(u, workerId, args[0].(*memory.Erew[int]), args[1].([]*memory.Erew[int]))
		}

		x := memory.AllocateMemoryWithData[int](m, elem)
		array := memory.AllocateMemorySlice[int](m, workers)

		m.Run(lambda, x, array)
		actual := memory.GetDataSlice(m, array)
		assert.Equal(t, expected, actual)
	})
}
