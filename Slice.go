package main

import "fmt"

func main() {
	var months = [...]string{
		"Jan",
		"Feb",
		"Mar",
		"Apr",
		"May",
		"Jun",
		"Jul",
		"Aug",
		"Sep",
		"Oct",
		"Nov",
		"Dec",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Wanhar")
	fmt.Println(slice3)
	slice3[1] = "Bukan December"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Wanhar"
	newSlice[1] = "Daeng Maro"

	fmt.Println(newSlice)
}
