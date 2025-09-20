package memory

import (
	"sync"
)

type emptyStruct struct{}

type Unique = Erew[emptyStruct]

// WorkerFunc TODO(kasalakhov): unique is crutch, think about change it
type WorkerFunc = func(unique *Unique, workerId int, args ...interface{})

type Manager struct {
	barrier     *Barrier
	workerCount int
}

func NewManager(workerCount int) *Manager {
	return &Manager{
		workerCount: workerCount,
		barrier:     NewBarrier(workerCount),
	}
}

func (m *Manager) Run(f WorkerFunc, args ...interface{}) {
	wg := &sync.WaitGroup{}
	wg.Add(m.workerCount)
	for workerId := range m.workerCount {
		go func() {
			defer wg.Done()
			f(AllocateMemory[emptyStruct](m), workerId, args...)
		}()
	}
	wg.Wait()
}

func (m *Manager) WorkerCount() int {
	return m.workerCount
}

func GetData[T any](mem *Erew[T]) T {
	return mem.getData()
}

func GetDataSlice[T any](m *Manager, mem []*Erew[T]) []T {
	length := len(mem)
	slice := make([]T, length)
	for i := range length {
		if mem[i].m != m {
			panic("mem runner incorrect, cannot getData")
		}
		slice[i] = mem[i].getData()
	}

	return slice
}
