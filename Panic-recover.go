package main

import "fmt"

func main() {
	runApp(true)
	fmt.Println("App still running")
}

func endApp() {
	fmt.Println("End application")
	message := recover()
	if message != nil {
		fmt.Println("Error with message: ", message)
	}
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Error application")
	}
	fmt.Println("Start application")
}
