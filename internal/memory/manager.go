package memory

import (
	"sync"
	"sync/atomic"
)

const MaxStepCount = 1_000_000

type EmptyStruct struct{}

type Unique = Erew[EmptyStruct]

// WorkerFunc TODO(kasalakhov): unique is crutch, think about change it
type WorkerFunc = func(unique *Unique, workerId int, args ...interface{})

type Manager struct {
	// TODO(kasalakhov): stupid solution, think about changing sync.map to another struct
	stepToCount map[int32]*atomic.Int32
	step        *atomic.Int32
	workerCount int
}

func NewManager(workerCount int) *Manager {
	stepToCount := make(map[int32]*atomic.Int32, MaxStepCount)
	for i := range MaxStepCount {
		stepToCount[int32(i)] = &atomic.Int32{}
		stepToCount[int32(i)].Store(int32(workerCount))
	}

	return &Manager{
		stepToCount: stepToCount,
		workerCount: workerCount,
		step:        &atomic.Int32{},
	}
}

func (m *Manager) Run(f WorkerFunc, args ...interface{}) {
	wg := &sync.WaitGroup{}
	wg.Add(m.workerCount)
	for workerId := range m.workerCount {
		go func() {
			defer wg.Done()
			f(AllocateMemory[EmptyStruct](m), workerId, args...)
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
