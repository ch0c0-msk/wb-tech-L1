package main

import "fmt"

type Human struct {
	age  int
	name string
}

func (h Human) Hello() {
	fmt.Printf("Hello, I am %s! I`m %d years old.\n", h.name, h.age)
}

type Action struct {
	Human
}

func main() {
	human := &Human{age: 18, name: "Michael"}
	action := &Action{Human: *human}
	action.Hello()
}
