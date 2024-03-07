package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, input <-chan interface{}) {
	defer wg.Done()
	// Если канал данных закрыт, завершаем горутину
	for data := range input {
		fmt.Println(data)
	}
}

func main() {
	var workersCount int
	fmt.Scan(&workersCount)
	data := make(chan interface{}, 1)
	var wg sync.WaitGroup

	// Запуск горутин
	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go worker(&wg, data)
	}

	// Канал в который запишется сигнал прерывания
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	for {
		select {
		// Если прерывание, то закрываем канал данных, ждем завершения горутин и выходим из программы
		case <-signalCh:
			close(data)
			wg.Wait()
			return
		// Генерация данных
		default:
			data <- time.Now().UnixNano()
			time.Sleep(1 * time.Second)
		}
	}
}
