package main

import "fmt"

func main() {
	// factorialLoop
	loop := factorialLoop(5)
	fmt.Println(loop)
	fmt.Println("============")
	fmt.Println(5 * 4 * 3 * 2 * 1)

	fmt.Println("============")

	// with recursive
	resultRecursive := factorialRecursive(5)
	fmt.Println(resultRecursive)
}

// without recursive
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

// with recursive
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}
