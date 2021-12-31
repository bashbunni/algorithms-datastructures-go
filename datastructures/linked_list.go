package datastructures

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func NewNode(value int) *Node {
	return &Node{value, nil}
}

func NewNodeWithNext(value int, next *Node) *Node {
	return &Node{value, next}
}

// Push returns the new head of the linked list
func (head *Node) Push(value int) *Node {
	return NewNodeWithNext(value, head)
}

func (head *Node) AddAfter(value int, before *Node) {
	curr := head
	for curr != nil {
		if curr.Val == before.Val {
			temp := before.Next
			node := NewNodeWithNext(value, temp)
			before.Next = node
		}
		curr = curr.Next
	}
}

func (head *Node) Append(value int) {
	curr := head
	if curr != nil {
		for curr.Next != nil {
			curr = curr.Next
		}
	}
	curr = NewNode(value)
}

func (head *Node) Pop() {
	curr := head
	for curr.Next.Next != nil {
		curr = curr.Next
	}
	curr.Next = nil
}

func BuildLinkedList() {
	a := NewNode(1)
	b := NewNode(2)
	c := NewNode(3)
	d := NewNode(4)

	a.Next = b
	b.Next = c
	c.Next = d
	fmt.Println(a.Next.Val)
	fmt.Println(b.Next.Val)
	fmt.Println(c.Next.Val)
	fmt.Println(a.Next.Next.Next.Val)
	a.Append(5)
	a.AddAfter(7, b)
	fmt.Println(b.Next)
	fmt.Println(c.Next)
	a.Pop()
	fmt.Println(c.Next)
}
