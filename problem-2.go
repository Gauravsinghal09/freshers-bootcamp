package main

import "fmt"

// Tree struct contains a root rune, left and right pointer to another Tree struct
type Tree struct{
	root rune
	left *Tree
	right *Tree
}

// Returns a new Tree
func newTree(root rune) *Tree{
	var tree = &Tree{root: root}
	tree.left = nil
	tree.right = nil
	return tree
}

// Preorder traversal
func preorder(tree *Tree){

	if tree != nil{
		fmt.Printf("%c ", tree.root)
		preorder(tree.left)
		preorder(tree.right)
	}
	return ;

}

// Postorder traversal
func postorder(tree *Tree){
	if tree != nil{
		postorder(tree.left)
		postorder(tree.right)
		fmt.Printf("%c ", tree.root)
	}
	return ;
}

func main(){
	var tree = newTree('+')
	tree.left = newTree('a')
	tree.right = newTree('-')
	tree.right.left = newTree('b')
	tree.right.right = newTree('c')

	preorder(tree)
	fmt.Println()
	postorder(tree)
	fmt.Println()


}
