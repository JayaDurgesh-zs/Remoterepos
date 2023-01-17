package main

import "fmt"

func Array() {
	var arr [5]int
	for i := 0; i < 5; i++ {
		fmt.Scanln(&arr[i])
		fmt.Printf("Element at %v is %v\n", i, arr[i])
	}
}
