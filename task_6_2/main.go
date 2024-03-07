package main

import (
	"fmt"
	"time"
)

func main() {
	in := make(chan interface{})
	go foo(in)
	in <- "data1"
	in <- "data2"
	in <- "data3"
	close(in)
	// Даем время на завершение горутины
	time.Sleep(time.Second)
}

// Остановка путем закрытия входного канала
func foo(in <-chan interface{}) {
	for data := range in {
		fmt.Printf("do something with input %s\n", data)
	}
	fmt.Println("finished")
}
