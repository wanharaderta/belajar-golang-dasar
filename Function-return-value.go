package main

import "fmt"

func main() {
	// single value
	result := getHello("Wanhar")
	fmt.Println(result)

	// multi value
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}

// Returning Single value
func getHello(name string) string {
	return "Hello," + name
}

// Returning Single value
func getFullName() (string, string) {
	return "Wanhar", "Daeng maro"
}
