package main

import (
	"fmt"
	"time"
)

// Оставляем такую же сигнатуру как в пакете time
func Sleep(dur time.Duration) {
	// Используем функцию After для получения канала, которые закроется через заданный интервал
	<-time.After(dur)
}

func main() {
	fmt.Println("Start")
	Sleep(2 * time.Second)
	fmt.Println("End")
}
