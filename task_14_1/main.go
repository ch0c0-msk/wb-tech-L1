package main

import "fmt"

func main() {
	test := []interface{}{25, "25", make(chan int), true}
	for _, val := range test {
		fmt.Println(getVarType(val))
	}
}

func getVarType(val interface{}) string {
	return fmt.Sprintf("%T", val)
}
