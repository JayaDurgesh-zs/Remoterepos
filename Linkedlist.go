package main

import "fmt"

type node struct {
	next *node
	val  int
}

func linkedlist(head *node, value int) {
	n := node{val: value}
	n1 := &n
	head.next = n1
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	h1 := node{val: arr[0]}
	head := &h1
	head1 := head
	for i := 1; i < len(arr); i++ {
		linkedlist(head, arr[i])
		head = head.next
	}
	for head1 != nil {
		if head1.next != nil {
			fmt.Print(head1.val, "->")
		} else {
			fmt.Println(head1.val)
		}
		head1 = head1.next
	}
}
