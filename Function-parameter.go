package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		sayHelloTo("Wanhar", "Daeng Maro")
	}
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}
