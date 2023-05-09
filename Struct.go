package main

import "fmt"

func main() {
	var user Customer
	user.name = "Wanhar"
	user.address = "Gowa"
	user.age = 27
	fmt.Println("User", user)

	customer := Customer{
		name:    "Wanhar",
		address: "Gowa",
		age:     27,
	}
	fmt.Println("Customer", customer)
}

type Customer struct {
	name    string
	address string
	age     int
}
