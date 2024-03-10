package main

import "fmt"

func main() {
	testSet1 := map[interface{}]bool{1: true, 2: true, 5: true, 7: true, 9: true}
	testSet2 := map[interface{}]bool{1: true, 3: true, 5: true, 7: true, 10: true}
	fmt.Println("set 1:")
	for key := range testSet1 {
		fmt.Printf("%v ", key)
	}
	fmt.Println()
	fmt.Println("set 2:")
	for key := range testSet2 {
		fmt.Printf("%v ", key)
	}
	fmt.Println()
	fmt.Println("intersection:")
	for key := range intersection(testSet1, testSet2) {
		fmt.Printf("%v ", key)
	}
	fmt.Println()
}

func intersection(firstSet, secondSet map[interface{}]bool) map[interface{}]bool {
	res := make(map[interface{}]bool)
	for elem1 := range firstSet {
		if _, isExist := secondSet[elem1]; isExist {
			res[elem1] = true
		}
	}
	return res
}
