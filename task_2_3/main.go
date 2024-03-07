package main

import (
	"fmt"
	"sync"
)

func main() {
	numList := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	res := make([]int, 5)

	// Результат вычисления отдельного числа записывается в общий слайс результатов на свое место по индексу
	for i, val := range numList {
		wg.Add(1)
		go func(res []int, i int, num int) {
			defer wg.Done()
			res[i] = num * num
		}(res, i, val)
	}
	wg.Wait()

	// Вывод результатов в исходном порядке
	for _, num := range res {
		fmt.Println(num)
	}
}
