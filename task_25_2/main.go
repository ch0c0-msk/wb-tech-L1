package main

import (
	"fmt"
	"time"
)

// То же самое, но через таймер
func Sleep(dur time.Duration) {
	<-time.NewTimer(dur).C
}

func main() {
	fmt.Println("Start")
	Sleep(2 * time.Second)
	fmt.Println("End")
}
