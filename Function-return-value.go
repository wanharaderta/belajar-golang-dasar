package main

import "fmt"

func main() {
	result := getHello("Wanhar")
	fmt.Println(result)
}

func getHello(name string) string {
	return "Hello," + name
}
