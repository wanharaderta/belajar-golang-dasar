package main

import "fmt"

func main() {
	// Pass by value
	address1 := Address{
		city:     "Jakarta",
		province: "DKI Jakarta",
		country:  "Indonesia",
	}
	address2 := address1

	address2.city = "Makassar"
	fmt.Println(address1)
	fmt.Println(address2)

	// Pointer
	fmt.Println("=====Pointer=====")
	address3 := &address1
	address3.city = "Bandung"
	fmt.Println(address1)
	fmt.Println(address3)

	// Operator
	fmt.Println("=====Operator=====")
	var address4 *Address = &address1

	*address4 = Address{
		city:     "Malang",
		province: "Jawa timur",
		country:  "Indonesia",
	}

	fmt.Println(address1)
	fmt.Println(address4)

	// Function New
	fmt.Println("=====Function New=====")
	address5 := new(Address)
	fmt.Println(address5)

	// Pointer function
	fmt.Println("=====Pointer function=====")
	var address6 = Address{
		city:     "Bogor",
		province: "Jawa Barat",
		country:  "",
	}
	ChangeCountryToIndonesian(&address6)
	fmt.Println(address6)
}

// Pointer function
func ChangeCountryToIndonesian(address *Address) {
	address.country = "Indonesian"
}

type Address struct {
	city, province, country string
}
