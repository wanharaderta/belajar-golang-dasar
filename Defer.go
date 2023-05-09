package main

import "fmt"

func main() {
	runApplication(5)
}

func logging() {
	fmt.Println("Finished Calling function ")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Result", result)
}
