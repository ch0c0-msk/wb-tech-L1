package main

import (
	"fmt"
	"sync"
)

func calcSquare(wg *sync.WaitGroup, num int, out chan<- int) {
	defer wg.Done()
	out <- (num * num)
}

func main() {
	numList := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	res := make(chan int, len(numList))

	// Вычисления в отдельной функции, а не в замыкании
	for _, val := range numList {
		wg.Add(1)
		go calcSquare(&wg, val, res)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	var sum int
	// Канал закроется и цикл прервется, когда отработают все горутины по рассчету квадрата
	for num := range res {
		sum += num
	}
	fmt.Println(sum)
}
