package main

import "fmt"

// Type Declaration
type Filter func(string) string

func main() {
	// Makes a function to be variable
	goodBye := getGoodBye
	fmt.Println(goodBye("Wanhar"))

	// Makes a function to be parameter
	sayHelloWithFilter("Wanhar", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
}

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}
