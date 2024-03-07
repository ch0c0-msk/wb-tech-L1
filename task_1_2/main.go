package main

import "fmt"

type Humaner interface {
	Hello()
}

type Human struct {
	age  int
	name string
}

func (h Human) Hello() {
	fmt.Printf("Hello, I am %s! I`m %d years old.\n", h.name, h.age)
}

type Action struct {
	Humaner
}

func main() {
	human := &Human{age: 18, name: "Michael"}
	action := &Action{Humaner: *human}
	action.Hello()
}
