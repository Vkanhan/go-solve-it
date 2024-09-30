package main 

import(
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode 
}

func invertBinaryTree(root *TreeNode) *TreeNode {
	//base case 
	if root == nil {
		return nil
	}

	//swap the left and right child of the root
	root.Left, root.Right = root.Right, root.Left 

	//recursively invert the left tree
	invertBinaryTree(root.Left)

	//recursively invert the right tree
	invertBinaryTree(root.Right)

	return root 
}

func printTree(root *TreeNode) {
	if root == nil {
		return 
	}

	printTree(root.Left)
	fmt.Printf("%d", root.Val)
	printTree(root.Right)
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Println("Original tree (in-order):")
	printTree(root)
	fmt.Println()

	// Invert the binary tree
	invertBinaryTree(root)

	fmt.Println("Inverted tree (in-order):")
	printTree(root)
	fmt.Println()
}