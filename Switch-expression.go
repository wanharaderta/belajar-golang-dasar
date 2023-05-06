package main

import "fmt"

func main() {
	name := "Wanhar"

	switch name {
	case "Wanhar":
		fmt.Println("Hi, Wanhar")
	case "Maro":
		fmt.Println("Maro")
	default:
		fmt.Println("Unknown")
	}

	// Short Statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Name too long")
	case false:
		fmt.Println("Name too short")
	}

	// Swtich without condition
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Name too long")
	case length > 5:
		fmt.Println("Quite a long name")
	default:
		fmt.Println("Name too short")
	}
}
