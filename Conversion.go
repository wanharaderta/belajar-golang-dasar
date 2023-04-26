package main

import "fmt"

func main() {
	var value23 int32 = 129
	var value62 int64 = int64(value23)
	var value8 int8 = int8(value23)

	fmt.Println(value23)
	fmt.Println(value62)
	fmt.Println(value8)

	var name = "Wanhar"
	var e byte = name[0]
	var eString string = string(e)

	fmt.Println(name)
	fmt.Println(eString)
}
