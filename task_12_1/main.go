package main

import (
	"fmt"
)

func main() {
	// Строго говоря, данный список является мультимножеством, создадим для него собственное множество, включающее все элементы списка по одному разу
	testList := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println("Set:")
	for elem := range initSet(testList) {
		fmt.Printf("%s ", elem)
	}
	fmt.Println()
}

func initSet(list []string) map[string]bool {
	res := make(map[string]bool)
	for _, elem := range list {
		res[elem] = true
	}
	return res
}
