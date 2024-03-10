package main

import "fmt"

func main() {
	a, b := 10, 25
	ptrA, ptrB := &a, &b
	fmt.Printf("a: %d, b: %d\n", a, b)
	swap(ptrA, ptrB)
	fmt.Printf("a: %d, b: %d\n", a, b)
}

// Меняем путем смены указателей
func swap(a, b *int) {
	*a, *b = *b, *a
}
