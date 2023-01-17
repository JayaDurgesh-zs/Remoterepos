package main

import (
	"fmt"
	"sort"
)

type P1 struct {
	name string
	age  int
}
type People struct {
	sli []P1
}

func main() {
	a1 := P1{"durgesh", 20}
	a2 := P1{"kumar", 22}
	a3 := P1{"jaya", 18}
	b := []P1{a1, a2, a3}
	t := People{b}
	fmt.Println(t)
	sort.Slice(b, func(x, y int) bool {
		return b[x].age < b[y].age
	})
	fmt.Println(b)
}
