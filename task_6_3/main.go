package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	in := make(chan interface{})
	ctx, cancel := context.WithCancel(context.Background())
	go foo(ctx, in)
	in <- "data1"
	in <- "data2"
	in <- "data3"
	// отменяем контекст
	cancel()
	// Даем время на завершение горутины
	time.Sleep(time.Second)
}

// Остановка путем отмены контекста
func foo(ctx context.Context, in <-chan interface{}) {
	for {
		select {
		case data := <-in:
			fmt.Printf("do something with input %s\n", data)
		case <-ctx.Done():
			fmt.Println("finished")
			return
		}
	}
}
