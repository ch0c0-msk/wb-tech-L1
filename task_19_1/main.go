package main

import (
	"fmt"
	"strings"
)

func main() {
	test := "aaabbbccc"
	fmt.Println(reverse(test))
}

// Reverse с помощью string builder
func reverse(in string) string {
	var res strings.Builder
	for i := len(in) - 1; i >= 0; i-- {
		res.WriteByte(in[i])
	}
	return res.String()
}
