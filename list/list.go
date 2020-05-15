package list

import "fmt"

type Node struct {
	Value int
	Next *Node
	Prev *Node
}

var TestList = &Node{ Value: 1 }

func init() {
	TestList.Next = &Node{ Value: 2 }
	TestList.Next.Next = &Node{ Value: 3 }
	TestList.Next.Next.Next = &Node{ Value: 4 }
	TestList.Next.Next.Next.Next = &Node{ Value: 5 }
}

// return head
func NewList(v int) *Node {
	return &Node{ Value: v }
}

func (n *Node) Add(v int) {
	n.Next = &Node{ Value: v }
}

func (n *Node) Remove(v int) {

}

func (n *Node) Paint() {
	fmt.Print(n.Value, "->")
	for n.Next != nil {
		n = n.Next
		fmt.Print(n.Value, "->")
	}
}