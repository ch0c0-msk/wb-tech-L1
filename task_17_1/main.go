package main

import (
	"fmt"
)

func binSearch(arr []int, key int) int {
	l := 0
	r := len(arr) - 1
	for l <= r {
		mid := (l + r) / 2
		if key < arr[mid] {
			r = mid - 1
		} else if key > arr[mid] {
			l = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	sortTest := []int{1, 2, 4, 9, 10, 10, 12, 15, 17}
	fmt.Println(binSearch(sortTest, 4))
}
