package main

import "fmt"

func main() {
	total := sumAll(10, 10, 10)
	fmt.Println(total)

	// convert slice to varargs
	slice := []int{10, 20, 30, 40, 50}
	total2 := sumAll(slice...)
	fmt.Println(total2)
}

// variable arguments (varargs)
func sumAll(number ...int) int {
	total := 0
	for _, value := range number {
		total += value
	}
	return total
}
