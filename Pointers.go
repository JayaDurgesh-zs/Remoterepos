package main

import "fmt"

func Pointers() {
	i := 42
	p := &i
	fmt.Println(*p)
	*p = 32
	fmt.Println(i)
}
