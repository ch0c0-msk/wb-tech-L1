package main

import (
	"fmt"
	"sync"
)

type SyncCounter struct {
	mu      sync.Mutex
	counter int
}

func (sc *SyncCounter) Inc() {
	sc.mu.Lock()
	sc.counter++
	sc.mu.Unlock()
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
