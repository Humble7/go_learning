package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) PrintValue() {
	fmt.Println(node.Value)
}

func (node *Node) SetValue(value int) {
	if node == nil {
        fmt.Println("Current Node is Nil!")
		return
	}
	node.Value = value
}


