package main

import "fmt"

func main() {
	a, b := float64(1<<511), float64(1<<511)
	fmt.Println(add(a, b))
	fmt.Println(sub(a, b))
	fmt.Println(devide(a, b))
	fmt.Println(multiply(a, b))
}

// Использование float64 с точностью до 15 цифр и порядка 308
func add(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}

func devide(a, b float64) float64 {
	return a / b
}

func multiply(a, b float64) float64 {
	return a * b
}
