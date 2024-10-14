package main

import (
	"fmt"
)

// ListNode represents a node in a singly linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// MergeTwoList merges two sorted linked lists into a single sorted linked list.
// It returns the head of the merged list.
func MergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	// Base cases: if either list is nil, return the other list
	if list1 == nil && list2 != nil {
		return list2
	}
	if list1 != nil && list2 == nil {
		return list1
	}
	if list1 == nil && list2 == nil {
		return nil
	}

	newNode := new(ListNode)

	if list1.Val >= list2.Val {
		newNode.Val = list2.Val
		list2 = list2.Next
		newNode.Next = MergeTwoList(list1, list2)
	} else {
		newNode.Val = list1.Val
		list1 = list1.Next
		newNode.Next = MergeTwoList(list1, list2)
	}

	return newNode
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	fmt.Println("Merge two sorted lists")

	list1 := &ListNode{Val: 5, Next: &ListNode{Val: 8, Next: &ListNode{Val: 11, Next: nil}}}
	list2 := &ListNode{Val: 6, Next: &ListNode{Val: 8, Next: &ListNode{Val: 12, Next: nil}}}

	fmt.Println("Merged List:")
	printList(MergeTwoList(list1, list2))
}
