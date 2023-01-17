package main

import "fmt"

func Structures() {
	type Address struct {
		Name    string
		PinCode int
	}
	fmt.Println(Address{"durgesh", 533101})
}
