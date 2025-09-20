package memory

import (
	"sync/atomic"
)

type Erew[T any] struct {
	locked atomic.Bool
	m      *Manager
	data   T
}

func AllocateMemory[T any](m *Manager) *Erew[T] {
	return &Erew[T]{
		m: m,
	}
}

func AllocateMemorySlice[T any](r *Manager, size int) []*Erew[T] {
	result := make([]*Erew[T], size)
	for i := range size {
		result[i] = AllocateMemory[T](r)
	}

	return result
}

func AllocateMemoryWithData[T any](r *Manager, data T) *Erew[T] {
	return &Erew[T]{
		m:    r,
		data: data,
	}
}

// Read blocks thread until N threads summary will call Read(), Write() or Skip()
func (e *Erew[T]) Read() T {
	if !e.locked.CompareAndSwap(false, true) {
		panic("Repeated memory access in one point in time by reading")
	}

	result := e.data
	e.m.barrier.Wait()
	e.locked.Store(false)
	e.m.barrier.Wait()

	return result
}

func (e *Erew[T]) Write(newData T) {
	if !e.locked.CompareAndSwap(false, true) {
		panic("Repeated memory access in one point in time by writing")
	}

	e.data = newData
	e.m.barrier.Wait()
	e.locked.Store(false)
	e.m.barrier.Wait()
}

func (e *Erew[EmptyStruct]) Skip(times int) {
	for range times {
		e.m.barrier.Wait()
		e.m.barrier.Wait()
	}
}

func (e *Erew[T]) getData() T {
	return e.data
}
