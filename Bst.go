package main

//
//import "fmt"
//
//type node struct {
//	left  *node
//	right *node
//	val   int
//}
//
//func CreateTree(root *node, nod *node) {
//	temp := root
//
//	if nod.val < temp.val {
//		if temp.left != nil {
//			CreateTree(temp.left, nod)
//		} else {
//			temp.left = nod
//
//		}
//	} else {
//		if temp.right != nil {
//			CreateTree(temp.right, nod)
//		} else {
//			temp.right = nod
//
//		}
//	}
//
//}
//
//func Preorder(root *node) {
//	if root != nil {
//		fmt.Println(root.val)
//		Preorder(root.left)
//		Preorder(root.right)
//	}
//}
//
//func main() {
//	arr := []int{4, 5, 3, 7, 2, 6, 9}
//	n := node{val: arr[0]}
//	root := &n
//	//fmt.Println(root.val)
//	//root1 := root
//	for i := 1; i < len(arr); i++ {
//		n1 := node{val: arr[i]}
//		nod := &n1
//		CreateTree(root, nod)
//	}
//	Preorder(root)
//
//}
