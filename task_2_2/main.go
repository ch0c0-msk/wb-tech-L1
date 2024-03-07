package main

import "fmt"

func main() {
	numList := [5]int{2, 4, 6, 8, 10}
	res := make(chan int, len(numList))

	for _, val := range numList {
		go func(num int, res chan<- int) {
			res <- (num * num)
		}(val, res)
	}

	// Берем первые 5 значений из канала
	for i := 0; i < len(numList); i++ {
		fmt.Println(<-res)
	}
}
