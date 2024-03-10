package main

import (
	"fmt"
)

func main() {
	test := "aaabbbccc"
	fmt.Println(reverse(test))
}

// Reverse с помощью преобразования в срез байтов
func reverse(in string) string {
	bytes := []byte(in)
	for i := 0; i < (len(in) / 2); i++ {
		bytes[i], bytes[len(in)-i-1] = bytes[len(in)-i-1], bytes[i]
	}
	return string(bytes)
}
