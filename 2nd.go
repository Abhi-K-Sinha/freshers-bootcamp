package main

import (
	"fmt"
)

type node struct {
	val   string
	left  *node
	right *node
}

func buildtree(exp string) *node {
	var arr []*node
	for _, x := range exp {
		if x >= '0' && x <= '9' {
			arr = append(arr, &node{val: string(x)})
		} else {
			tmp := &node{val: string(x)}
			tmp.left = arr[len(arr)-2]
			tmp.right = arr[len(arr)-1]
			arr = arr[:len(arr)-2]
			arr = append(arr, tmp)
		}
	}
	return arr[0]

}
func inorder(root *node) {
	if root == nil {
		return
	}
	inorder(root.left)
	fmt.Printf("%v ", root.val)
	inorder(root.right)
}
func preorder(root *node) {
	if root == nil {
		return
	}
	fmt.Printf("%v ", root.val)
	preorder(root.left)
	preorder(root.right)
}
func postorder(root *node) {
	if root == nil {
		return
	}
	postorder(root.left)
	postorder(root.right)
	fmt.Printf("%v ", root.val)
}
func main() {
	exp := "123*+4/"
	root := buildtree(exp)
	fmt.Println("Inorder:")
	inorder(root)
	fmt.Println("\nPreorder:")
	preorder(root)
	fmt.Println("\nPostorder:")
	postorder(root)

}
