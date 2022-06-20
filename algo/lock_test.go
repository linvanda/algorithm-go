package algo

import (
	"sync"
	"testing"
)

var mtx sync.Mutex

func BenchmarkLock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mtx.Lock()
		mtx.Unlock()
	}
}
