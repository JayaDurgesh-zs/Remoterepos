package main

import "fmt"

type node struct {
	next *node
	prev *node
	val  int
}

func Doublyllist(h *node, value int) {
	v := node{val: value}
	nod := &v
	h.next = nod
	nod.prev = h
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	n := node{val: arr[0]}
	head := &n
	head1 := head
	for i := 1; i < len(arr); i++ {
		Doublyllist(head, arr[i])
		head = head.next
	}
	for head1 != nil {
		if head1.next != nil {
			fmt.Print(head1.val, "<->")
		} else {
			fmt.Print(head1.val)
		}
		head1 = head1.next
	}
}
