package main

import (
	"fmt"
	"strings"
)

func main() {
	test := "snow dog sun"
	fmt.Println(reverseWords(test))
}

// Получаем из строки срез слов и просто переворачиваем срез
func reverseWords(in string) string {
	words := strings.Split(in, " ")
	for i := 0; i < (len(words) / 2); i++ {
		words[i], words[len(words)-i-1] = words[len(words)-i-1], words[i]
	}
	return strings.Join(words, " ")
}
