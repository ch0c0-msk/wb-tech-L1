package main

import "fmt"

const number int64 = -4

func main() {
	var bitNumber, value int
	fmt.Scan(&bitNumber)
	fmt.Scan(&value)
	fmt.Printf("Old value: %b\n", number)
	newNumber := setBit(number, bitNumber, value)
	fmt.Printf("New value: %b\n", newNumber)
}

func setBit(number int64, bitIndex int, value int) int64 {
	res := number
	isNegative := false
	// Если число отрицательное, получаем его прямой код и работем с ним. Пример: -4 -> 4
	if number < 0 {
		res = getTwosComplement(number)
		isNegative = true
	}

	// Меняем нужный бит
	if value == 0 {
		res &= ^(1 << bitIndex)
	} else {
		res |= 1 << bitIndex
	}

	// Если число отрицательное, вернуть тоже надо отрицательное после смены нужного бита
	if isNegative {
		return -res
	}
	return res
}

// Получение доп. кода при работе с отрицательными числами
func getTwosComplement(number int64) int64 {
	return ^number + 1
}
