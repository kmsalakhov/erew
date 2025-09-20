package memory

import (
	"sync"
)

type Barrier struct {
	cond       *sync.Cond
	count      int
	generation int
	threshold  int
}

// Wait barrier with active waiting can be more effective on big count of goroutines, see benchmark bellow
func (b *Barrier) Wait() {
	b.cond.L.Lock()
	defer b.cond.L.Unlock()

	b.count++
	generation := b.generation

	if b.count < b.threshold {
		for generation == b.generation {
			b.cond.Wait()
		}
	} else {
		b.count = 0
		b.generation++
		b.cond.Broadcast()
	}
}

func NewBarrier(threshold int) *Barrier {
	return &Barrier{
		cond:       sync.NewCond(&sync.Mutex{}),
		count:      0,
		generation: 0,
		threshold:  threshold,
	}
}

/*
goos: darwin
goarch: arm64
pkg: erew/cmd
cpu: Apple M3 Pro
BenchmarkSetXToArrActiveWaitingSmall       	   97380	     12382 ns/op	     952 B/op	      23 allocs/op
BenchmarkSetXToArrActiveWaitingSmall-2     	   41434	     28973 ns/op	     952 B/op	      23 allocs/op
BenchmarkSetXToArrActiveWaitingSmall-4     	   16617	     73030 ns/op	     952 B/op	      23 allocs/op
BenchmarkSetXToArrActiveWaitingSmall-8     	   10135	    127993 ns/op	     952 B/op	      23 allocs/op
BenchmarkSetXToArrActiveWaitingSmall-16    	    3549	    646035 ns/op	     980 B/op	      23 allocs/op
BenchmarkSetXToArrSmall                    	   95733	     12575 ns/op	     952 B/op	      23 allocs/op
BenchmarkSetXToArrSmall-2                  	   48916	     24487 ns/op	     952 B/op	      23 allocs/op
BenchmarkSetXToArrSmall-4                  	   37005	     32338 ns/op	     952 B/op	      23 allocs/op
BenchmarkSetXToArrSmall-8                  	   34096	     34882 ns/op	     952 B/op	      23 allocs/op
BenchmarkSetXToArrSmall-16                 	   31850	     37634 ns/op	     952 B/op	      23 allocs/op
BenchmarkSetXToArrActiveWaitingBig         	      24	  43717842 ns/op	  880072 B/op	   20003 allocs/op
BenchmarkSetXToArrActiveWaitingBig-2       	      22	  52074710 ns/op	  880076 B/op	   20003 allocs/op
BenchmarkSetXToArrActiveWaitingBig-4       	      16	  63171164 ns/op	  880072 B/op	   20003 allocs/op
BenchmarkSetXToArrActiveWaitingBig-8       	      18	  59328567 ns/op	  883830 B/op	   20011 allocs/op
BenchmarkSetXToArrActiveWaitingBig-16      	      13	 104981635 ns/op	  893029 B/op	   20031 allocs/op
BenchmarkSetXToArrBig                      	      22	  52667100 ns/op	  923359 B/op	   20453 allocs/op
BenchmarkSetXToArrBig-2                    	      15	  79443589 ns/op	  943630 B/op	   20665 allocs/op
BenchmarkSetXToArrBig-4                    	       8	 125485729 ns/op	  997840 B/op	   21229 allocs/op
BenchmarkSetXToArrBig-8                    	       8	 136051912 ns/op	  994120 B/op	   21191 allocs/op
BenchmarkSetXToArrBig-16                   	       8	 136665271 ns/op	  990128 B/op	   21129 allocs/op
*/
