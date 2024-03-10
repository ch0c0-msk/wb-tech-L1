package main

import (
	"fmt"
	"strings"
)

var justString string

// Функция createHugeString судя по названию возвращает длинную строку, однако ее размер не гарантированно больше 100, поэтому необходимо проверять длину
func someFunc() {
	v := createHugeString(1 << 10)
	if len(v) >= 100 {
		justString = v[:100]
	} else {
		justString = v
	}
}

// Какая-то реализация
func createHugeString(size int) string {
	var res strings.Builder
	bytes := make([]byte, size)
	for i := range bytes {
		bytes[i] = 'a'
	}
	res.Write(bytes)
	return res.String()
}

func main() {
	someFunc()
	fmt.Println(justString)
}
