package main 

import(
	"fmt"
)

//Singly linked list
type ListNode struct {
	Val int
	Next *ListNode
}

func middleOfLinkedList(head *ListNode) *ListNode {
	//Initialize two pointers
	slow, fast := head, head

	//Move fast pointers two steps and slow pointer one step
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	//Return the middle node
	return slow
}

func main() {
	// Creating a linked list 1 -> 2 -> 3 -> 4 -> 5
    head := &ListNode{Val: 1}
    head.Next = &ListNode{Val: 2}
    head.Next.Next = &ListNode{Val: 3}
    head.Next.Next.Next = &ListNode{Val: 4}
    head.Next.Next.Next.Next = &ListNode{Val: 5}

    middle := middleOfLinkedList(head)
    fmt.Println(middle.Val)
}