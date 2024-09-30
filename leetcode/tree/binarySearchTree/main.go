package main 

import(
	"fmt"
)

//TreeNode represents each node of the binary search tree
type TreeNode struct {
	Val int 
	Left *TreeNode
	Right *TreeNode
}

//insert a value into BST
func insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if val < root.Val {
		root.Left = insert(root.Left, val)
	} else if val > root.Val {
		root.Right = insert(root.Right, val)
	}
	return root 
} 

//Search a value in BST
func search(root *TreeNode, val int) bool {
	if root == nil {
		return false 
	}

	if root.Val == val {
		return true 
	}

	if val < root.Val {
		return search(root.Left, val)
	}
	return search(root.Right, val)
}

//Find the minimum value node in the BST
func findMin(root *TreeNode) *TreeNode {
	current := root 
	for current.Left != nil {
		current = current.Left
	}
	return current
}

//Delete a value from the BST
func delete(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil  
	}

	//Find the node to delete
	if val < root.Val {
		root.Left = delete(root.Left, val)
	} else if val > root.Val {
		root.Right = delete(root.Right, val)
	} else {
		//Node found
		//Case 1: Node with no children
		if root.Left == nil && root.Right == nil {
			return nil 
		}
		//Case 2: Node with only 1 child
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}
		
		//Case 3: Node with two children
		//Find the smallest value in the right subtree (in-order successor)
		minRight := findMin(root.Right)

		// Replace the value of the node to be deleted with the in-order successor's value
		root.Val = minRight.Val

		// Delete the in-order successor in the right subtree
		root.Right = delete(root.Right, minRight.Val)
	}
	return root
}

//Inorder traversal of the tree to print the value in sorted order
func inorderTraversal(root *TreeNode) {
	if root == nil {
		return 
	}

	inorderTraversal(root.Left)
	fmt.Printf("%d", root.Val)
	inorderTraversal(root.Right)
}

func main() {
	var root *TreeNode

	// Insert values into the BST
	values := []int{8, 3, 10, 1, 6, 14, 4, 7, 13}
	for _, val := range values {
		root = insert(root, val)
	}

	// Print the BST using in-order traversal
	fmt.Println("In-order traversal of the BST:")
	inorderTraversal(root)
	fmt.Println()

	// Search for a value
	val := 6
	if search(root, val) {
		fmt.Printf("Value %d found in the BST\n", val)
	} else {
		fmt.Printf("Value %d not found in the BST\n", val)
	}

	// Delete a value from the BST
	deleteVal := 3
	fmt.Printf("Deleting value %d from the BST...\n", deleteVal)
	root = delete(root, deleteVal)

	// Print the BST after deletion
	fmt.Println("In-order traversal after deletion:")
	inorderTraversal(root)
	fmt.Println()
}