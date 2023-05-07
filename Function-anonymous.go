package main

import "fmt"

type Blacklist func(string) bool

func main() {
	blacklist := func(name string) bool { // anonymous function
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("wanhar", blacklist)
}

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are already blacklisted", name)
	} else {
		fmt.Println("Welcome to", name)
	}
}
