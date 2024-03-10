package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, b := big.NewFloat(1<<300), big.NewFloat(1<<511)
	fmt.Println(add(a, b))
	fmt.Println(sub(a, b))
	fmt.Println(devide(a, b))
	fmt.Println(multiply(a, b))
}

// Использование пакета big для работы с большими числами
func add(a, b *big.Float) *big.Float {
	return big.NewFloat(0).Add(a, b)
}

func sub(a, b *big.Float) *big.Float {
	return big.NewFloat(0).Sub(a, b)
}

func devide(a, b *big.Float) *big.Float {
	return big.NewFloat(0).Quo(a, b)
}

func multiply(a, b *big.Float) *big.Float {
	return big.NewFloat(0).Mul(a, b)
}
