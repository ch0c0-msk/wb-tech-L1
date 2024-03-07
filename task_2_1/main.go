package main

import (
	"fmt"
	"sync"
)

func main() {
	numList := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	res := make(chan int, len(numList))

	for _, val := range numList {
		wg.Add(1)
		go func(num int, res chan<- int) {
			defer wg.Done()
			res <- (num * num)
		}(val, res)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	// Канал закроется и цикл прервется, когда отработают все горутины по рассчету квадрата
	for num := range res {
		fmt.Printf("%d ", num)
		fmt.Println()
	}
}
