package main

import (
	"fmt"
)

// Define the ListNode structure
type ListNode struct {
	Val  int
	Next *ListNode
}

// Iterative method to reverse the list
func reverseListIterative(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

// Recursive method to reverse the list
func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// Print the list
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	fmt.Println("Original List:")
	printList(head)

	iterReversed := reverseListIterative(head)
	fmt.Println("Reversed Iteratively:")
	printList(iterReversed)

	head = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	recursiveReversed := reverseListRecursive(head)
	fmt.Println("Reversed Recursively:")
	printList(recursiveReversed)
}
