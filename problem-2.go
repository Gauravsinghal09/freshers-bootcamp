package main

import (
	"fmt"
)

// Tree struct contains a root rune, left and right pointer to another Tree struct
type Tree struct {
	root  rune
	left  *Tree
	right *Tree
}

// newTree function returns a new Tree struct
func newTree(root rune) *Tree {
	var tree = &Tree{root: root}
	tree.left = nil
	tree.right = nil
	return tree
}

// preorder traversal
func (tree *Tree) preorder() {

	if tree != nil {
		fmt.Printf("%c ", tree.root)
		tree.left.preorder()
		tree.right.preorder()
	}
}

// postorder traversal
func (tree *Tree) postorder() {
	if tree != nil {
		tree.left.postorder()
		tree.right.postorder()
		fmt.Printf("%c ", tree.root)
	}
}

func main() {

	var tree = newTree('+')
	tree.left = newTree('a')
	tree.right = newTree('-')
	tree.right.left = newTree('b')
	tree.right.right = newTree('c')

	fmt.Printf("Preorder: ")
	tree.preorder()
	fmt.Println()
	fmt.Printf("Postorder: ")
	tree.postorder()
	fmt.Println()

}
