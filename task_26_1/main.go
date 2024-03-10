package main

import "fmt"

func main() {
	test := []string{"abcd", "abCdefAaf", "aabcd", "aA"}
	for _, str := range test {
		fmt.Println(isUniq(str))
	}
}

func isUniq(in string) bool {
	uniqMap := make(map[rune]bool)
	isUniq := true
	for _, r := range in {
		if _, isExist := uniqMap[r]; isExist {
			isUniq = false
			break
		}
		uniqMap[r] = true
	}
	return isUniq
}
