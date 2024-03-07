package main

import (
	"fmt"
	"strconv"
	"sync"
)

type SyncMap struct {
	mu        sync.Mutex
	normalMap map[string]string
}

func NewSyncMap() *SyncMap {
	return &SyncMap{normalMap: make(map[string]string)}
}

func (s *SyncMap) Add(key, value string) {
	s.mu.Lock()
	s.normalMap[key] = value
	s.mu.Unlock()
}

func main() {
	syncMap := NewSyncMap()
	n := 10
	var wg sync.WaitGroup
	// Каждая горутина запишет в свой ключ свое значение
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			syncMap.Add(strconv.Itoa(i), "somevalue"+strconv.Itoa(i))
		}(i)
	}
	wg.Wait()
	fmt.Println(syncMap.normalMap)

	// Каждая горутина будет писать в ключ "9" свое значение
	syncMap = NewSyncMap()
	lastKey := 9
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			syncMap.Add(strconv.Itoa(lastKey), "somevalue"+strconv.Itoa(i))
		}(i)
	}
	wg.Wait()
	fmt.Println(syncMap.normalMap)
}
