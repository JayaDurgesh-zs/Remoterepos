package main

import "fmt"

func Switch(x, y int, op string) {
	switch op {
	case "+":
		fmt.Println(x + y)
	case "-":
		fmt.Println(x - y)
	case "*":
		fmt.Println(x * y)
	case "/":
		if x < y {
			fmt.Println(float64(x % y))
		} else {
			fmt.Println(x / y)
		}
	}
}
