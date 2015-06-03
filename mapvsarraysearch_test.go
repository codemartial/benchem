package benchem_test

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

const (
	N = 2
)

type E struct {
	k string
	v int32
}

var m = make(map[string]int32, N)
var indices = make([]string, N)
var arr = make([]E, N)
var mu = sync.Mutex{}

func prepindices() {
	for i := 0; i < N; i++ {
		indices = append(indices, fmt.Sprintf("%d", rand.Intn(N)))
	}
}

func prepmap() {
	for i := 0; i < N; i++ {
		key := fmt.Sprintf("%d", i)
		m[key] = int32(i)
	}
}

func preparray() {
	for i := 0; i < N; i++ {
		arr[i] = E{fmt.Sprintf("%d", i), int32(i)}
	}
}

func TestMapAcess(t *testing.T) {
	var wg = sync.WaitGroup{}
	go func() {
		wg.Add(1)
		for i := 0; i < 1000000; i++ {
			_,_ = m[indices[0]]
		}
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		for i := 0; i < 1000000; i++ {
			mu.Lock()
			m[indices[1]] = 1
			mu.Unlock()
		}
		wg.Done()
	}()
	wg.Wait()
	t.Log("")
}

func BenchmarkMap(b *testing.B) {
	prepindices()
	preparray()
	prepmap()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = m[indices[i%N]]
	}
}

func BenchmarkArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range arr {
			if v.k == indices[i%N] {
				break
			}
		}
	}
}
