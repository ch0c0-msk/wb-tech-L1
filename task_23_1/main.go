package main

import "fmt"

func main() {
	test := []int{1, 2, 3, 4, 5}
	fmt.Println(test)
	keyToRemove := 3
	// Удаляем с помощью оператора среза и функции append
	test = append(test[:keyToRemove], test[(keyToRemove+1):]...)
	fmt.Println(test)
}
