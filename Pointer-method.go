package main

import "fmt"

func main() {
	maro := Man{name: "maro"}
	maro.Married()

	fmt.Println(maro.name)
}

type Man struct {
	name string
}

func (man *Man) Married() {
	man.name = "Mr. " + man.name
}
