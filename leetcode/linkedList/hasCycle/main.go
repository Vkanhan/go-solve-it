package main 

import(
	"fmt"
)

type ListNode struct {
	Val int 
	Next *ListNode
}

func hasCycle(head *ListNode) bool{
	slow, fast := head, head 

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true 
		}
	}
	return false 
}

func main() {
	head := &ListNode{Val: 1}
    second := &ListNode{Val: 2}
    third := &ListNode{Val: 3}
    head.Next = second
    second.Next = third
    third.Next = second // Creates a cycle

    fmt.Println(hasCycle(head))
}