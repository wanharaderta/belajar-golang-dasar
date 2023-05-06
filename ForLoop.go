package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Loop to", counter)
		counter++
	}

	// For Statement
	for counter2 := 1; counter2 <= 10; counter2++ {
		fmt.Println("Loop to", counter2)
	}

	// Loop with indexed
	slice := []string{"Wanhar", "Daeng", "maro"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// For Ranges
	for index, value := range slice {
		fmt.Println("Index", index, "=", value)
	}
}
