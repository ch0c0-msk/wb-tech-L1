package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	in := make(chan int, 5)
	out := make(chan int, 5)
	// Запускаем воркер
	go worker(in, out)

	// Подаем данный на входной канал
	for _, num := range arr {
		in <- num
	}
	close(in)

	// Выводим результаты на консоль
	for data := range out {
		fmt.Println(data)
	}
}

func worker(in <-chan int, out chan<- int) {
	for data := range in {
		out <- (data * data)
	}
	close(out)
}
