package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	writeAndWriteCount = 10
)

func TestRunner_Run(t *testing.T) {
	t.Run("write-write", func(t *testing.T) {
		m := NewManager(writeAndWriteCount)

		lambda := func(u *Unique, workerId int, args ...interface{}) {
			writeWrite(u, workerId, args[0].([]*Erew[int]))
		}

		memSlice := AllocateMemorySlice[int](m, writeAndWriteCount)
		m.Run(lambda, memSlice)

		slice := GetDataSlice(m, memSlice)
		for i := range writeAndWriteCount {
			assert.Equal(t, slice[i], writeAndWriteCount-i-1)
		}
	})

	t.Run("write-read-write", func(t *testing.T) {
		r := NewManager(writeAndWriteCount)

		lambda := func(u *Unique, workerId int, args ...interface{}) {
			writeReadWrite(u, workerId, args[0].([]*Erew[int]))
		}

		memSlice := AllocateMemorySlice[int](r, writeAndWriteCount)
		r.Run(lambda, memSlice)

		slice := GetDataSlice(r, memSlice)
		for i := range writeAndWriteCount {
			assert.Equal(t, slice[i], writeAndWriteCount-i-1)
		}
	})
}

func writeWrite(u *Unique, workerId int, array []*Erew[int]) {
	array[workerId].Write(workerId)
	array[writeAndWriteCount-workerId-1].Write(workerId)
}

func writeReadWrite(u *Unique, workerId int, array []*Erew[int]) {
	array[workerId].Write(workerId)
	_ = array[workerId].Read()
	array[writeAndWriteCount-workerId-1].Write(workerId)
}
