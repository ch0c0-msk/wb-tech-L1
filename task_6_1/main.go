package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	go foo(done)
	time.Sleep(3 * time.Second)
	done <- struct{}{}
	// Завершаем главный поток, когда завершится горутина
	<-done
}

// Остановка по дополнительному каналу
func foo(done chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("finished")
			done <- struct{}{}
			return
		default:
			fmt.Println("do something")
			time.Sleep(time.Second)
		}
	}
}
