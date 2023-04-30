package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "John",
		"address": "123 Main St",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	// Delete data from map
	delete(person, "title")
	fmt.Println(person)
}
