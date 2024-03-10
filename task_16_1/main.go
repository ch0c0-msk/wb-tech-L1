package main

import (
	"fmt"
	"math/rand"
)

func partition(l, r int, arr []int) (int, int) {
	pivotVal := arr[rand.Intn(r-l+1)+l]
	e, g := l, l
	for n := l; n <= r; n++ {
		if arr[n] < pivotVal {
			arr[g], arr[n] = arr[n], arr[g]
			arr[e], arr[g] = arr[g], arr[e]
			e++
			g++
		} else if arr[n] == pivotVal {
			arr[g], arr[n] = arr[n], arr[g]
			g++
		}
	}
	return e, g
}

func quickSort(low, high int, arr []int) {
	if low < high {
		pivotLow, pivotHigh := partition(low, high, arr)
		quickSort(low, pivotLow-1, arr)
		quickSort(pivotHigh, high, arr)
	}
}

func main() {
	test := []int{1, 5, 3, 7, 4, 15, 10, 9}
	quickSort(0, len(test)-1, test)
	for i := range test {
		fmt.Printf("%d ", test[i])
	}
	fmt.Println()
}
