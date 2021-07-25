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

// NewTree function returns a new Tree struct
func NewTree(root rune) *Tree {
	var tree = &Tree{root: root}
	tree.left = nil
	tree.right = nil
	return tree
}

// Preorder traversal
func (tree *Tree) Preorder() {

	if tree != nil {
		fmt.Printf("%c ", tree.root)
		tree.left.Preorder()
		tree.right.Preorder()
	}
}

// Postorder traversal
func (tree *Tree) Postorder() {
	if tree != nil {
		tree.left.Postorder()
		tree.right.Postorder()
		fmt.Printf("%c ", tree.root)
	}
}

func main() {
	var tree = NewTree('+')
	tree.left = NewTree('a')
	tree.right = NewTree('-')
	tree.right.left = NewTree('b')
	tree.right.right = NewTree('c')

	fmt.Printf("Preorder: ")
	tree.Preorder()
	fmt.Println()
	fmt.Printf("Postorder: ")
	tree.Postorder()

	fmt.Println()

}
