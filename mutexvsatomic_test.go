package benchem_test

import (
	"sync"
	"sync/atomic"
	"testing"
)

var a int32

const (
	READS  = 10000
	WRITES = 10
)

var readrwmu = func(mu *sync.RWMutex) {
	for x := 0; x < READS; x++ {
		mu.RLock()
		_ = a
		mu.RUnlock()
	}
}

var writerwmu = func(mu *sync.RWMutex) {
	for x := 0; x < WRITES; x++ {
		mu.Lock()
		a = int32(x)
		mu.Unlock()
	}
}

var writemu = func(mu *sync.Mutex) {
	for x := 0; x < WRITES; x++ {
		mu.Lock()
		a = int32(x)
		mu.Unlock()
	}
}

var read = func() {
	for x := 0; x < READS; x++ {
		_ = atomic.LoadInt32(&a)
	}
}

var write = func() {
	for x := 0; x < WRITES; x++ {
		atomic.StoreInt32(&a, int32(x))
	}
}

var readuns = func() {
	for x := 0; x < READS; x++ {
		_ = a
	}
}
var writeuns = func() {
	for x := 0; x < WRITES; x++ {
		a = int32(x)
	}
}

func BenchmarkMutex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			wg := sync.WaitGroup{}
			mu := sync.RWMutex{}
			go func() {
				wg.Add(1)
				writerwmu(&mu)
				wg.Done()
			}()
			go func() {
				wg.Add(1)
				readrwmu(&mu)
				wg.Done()
			}()
			wg.Wait()
		}
	})

}

func BenchmarkAtomic(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			wg := sync.WaitGroup{}
			go func() {
				wg.Add(1)
				write()
				wg.Done()
			}()
			go func() {
				wg.Add(1)
				read()
				wg.Done()
			}()
			wg.Wait()
		}
	})

}

func BenchmarkUnsafeReads(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu := sync.Mutex{}
			wg := sync.WaitGroup{}
			go func() {
				wg.Add(1)
				writemu(&mu)
				wg.Done()
			}()
			go func() {
				wg.Add(1)
				readuns()
				wg.Done()
			}()
			wg.Wait()
		}
	})
}

func BenchmarkUnsafe(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			wg := sync.WaitGroup{}
			go func() {
				wg.Add(1)
				writeuns()
				wg.Done()
			}()
			go func() {
				wg.Add(1)
				readuns()
				wg.Done()
			}()
			wg.Wait()
		}
	})
}
