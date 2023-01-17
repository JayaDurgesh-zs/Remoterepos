package main

import "fmt"

func Loop() {
	var input int
fun:
	fmt.Println("You are not eligible to vote ")
	fmt.Println("Enter your age ")
	fmt.Scanln(&input)
	if input <= 17 {
		goto fun
	} else {
		fmt.Println("You can vote ")
	}
}
