package main

import "fmt"

func Interface(i interface{}) {
	fmt.Printf("%T\n", i)
	switch i.(type) {
	case int:
		fmt.Println("it is integer type")
	case string:
		fmt.Println("it is a string type")
	case struct{ name string }:
		fmt.Println("it is a struct type")
	default:
		fmt.Println("other than that")
	}
}
