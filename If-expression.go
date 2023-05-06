package main

import "fmt"

func main() {

	var name = "Wanhar"

	if name == "Wanhar" {
		fmt.Println("Hello Wanhar")
	} else if name == "Maro" {
		fmt.Println("Hello Maro")
	} else {
		fmt.Println("Hi, What your name?")
	}

	// Short Statement
	if length := len(name); length > 5 {
		fmt.Println("Name too Long")
	} else {
		fmt.Println("Name too short")
	}
}
