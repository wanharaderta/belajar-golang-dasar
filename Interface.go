package main

import "fmt"

func main() {
	wanhar := Person{
		name: "wanhar",
	}
	SayHello(wanhar)
}

type HasName interface {
	GetName() string
}

type Person struct {
	name string
}

func (person Person) GetName() string {
	return person.name
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}
