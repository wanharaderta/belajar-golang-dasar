package main

import "fmt"

func main() {
	a, b, c := getCompleteName()
	fmt.Println("Full Name: ", a, b, c)
}

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Wanhar"
	middleName = "Daeng"
	lastName = "Maro"
	return
}
