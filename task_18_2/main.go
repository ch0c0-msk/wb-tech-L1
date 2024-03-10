package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type SyncCounter struct {
	counter int64
}

// Конкуретная запись через атомарные операции
func (sc *SyncCounter) Inc() {
	atomic.AddInt64(&sc.counter, 1)
}

func main() {
	var sc SyncCounter
	var wg sync.WaitGroup
	n := 100
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sc.Inc()
		}()
	}
	wg.Wait()
	fmt.Printf("Excepted: %d; Actual: %d\n", n, sc.counter)
}
