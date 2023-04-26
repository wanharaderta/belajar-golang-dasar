package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpWanhar NoKTP = "123123123123123"
	var marriedStatus Married = true

	fmt.Println(noKtpWanhar)
	fmt.Println(marriedStatus)
}
