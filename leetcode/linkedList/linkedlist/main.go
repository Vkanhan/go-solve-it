package main

import (
	"fmt"
)

// Node represents a node in the linked list.
type Node struct {
	Value int
	Next  *Node
}

// LinkedList represents the linked list.
type LinkedList struct {
	Head *Node
}

// Add adds a new node with the given value to the end of the list.
func (l *LinkedList) Add(value int) {
	newNode := &Node{Value: value}
	if l.Head == nil {
		l.Head = newNode
		return
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// Delete removes the first occurrence of the given value from the list.
func (l *LinkedList) Delete(value int) {
	if l.Head == nil {
		return
	}

	if l.Head.Value == value {
		l.Head = l.Head.Next
		return
	}

	current := l.Head
	for current.Next != nil && current.Next.Value != value {
		current = current.Next
	}

	if current.Next != nil {
		current.Next = current.Next.Next
	}
}

// Print displays the values in the list.
func (l *LinkedList) Read() {
	current := l.Head
	for current != nil {
		fmt.Print(current.Value, " ")
		current = current.Next
	}
	fmt.Println()
}

func main() {
	l := &LinkedList{}

	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Read()

	l.Delete(2)
	l.Read() 

	l.Delete(1)
	l.Read() 

	l.Delete(3)
	l.Read() 
}
