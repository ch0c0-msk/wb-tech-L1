package main

import (
	"fmt"
	"time"
)

func main() {
	var n int
	fmt.Scan(&n)
	data := make(chan interface{})
	stop := make(chan struct{})

	go func(out chan<- interface{}, stop <-chan struct{}) {
		for {
			select {
			// При сигнале остановки, закрываем канал, тем самым останавливаем горутину приемник
			case <-stop:
				fmt.Println("End sender")
				close(out)
				return
			default:
				out <- time.Now().UnixNano()
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(data, stop)

	go func(in <-chan interface{}, stop chan<- struct{}) {
		for data := range in {
			fmt.Println(data)
		}
		fmt.Println("End receiver")
		// Уведомляем главный поток, что закончили работу
		stop <- struct{}{}
	}(data, stop)

	<-time.After(time.Second * time.Duration(n))
	stop <- struct{}{}
	<-stop
}
