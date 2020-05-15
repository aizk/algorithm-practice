package main

import (
	"fmt"
	"github.com/liunian1004/algorithm-practice/list"
)

func main() {
	RecursePrintList(list.TestList)
}

func RecursePrintList(n *list.Node) {
	if n.Next == nil {
		fmt.Print(n.Value)
		return
	}
	RecursePrintList(n.Next)
	fmt.Print(n.Value)
}
