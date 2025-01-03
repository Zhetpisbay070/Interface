package main

import "fmt"

type Node struct {
	N    int
	Next *Node
}

func printNode(node *Node) {
	for node != nil {
		node = node.Next
		fmt.Println(node.N)

	}

}

func main() {

	k := Node{
		N:    3,
		Next: nil,
	}
	l := Node{
		N:    2,
		Next: &k,
	}
	p := Node{
		N:    1,
		Next: &l,
	}
	printNode(&p)
}
