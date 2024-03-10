package main

import (
	"fmt"
	"reflect"
)

func main() {
	test := []interface{}{25, "25", make(chan int), true}
	for _, val := range test {
		fmt.Println(getVarType(val))
	}
}

func getVarType(val interface{}) reflect.Type {
	return reflect.TypeOf(val)
}
