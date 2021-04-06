package main

import (
	"fmt"
	"net_easy/tree"
)

// 通过组合的方式实现扩展
type myTreeNode struct {
	node *tree.Node
}

func (node *myTreeNode) postorderTraversal() {
    if node == nil || node.node == nil {
		return
	}

	left := myTreeNode{node.node.Left}
	left.postorderTraversal()
	right := myTreeNode{node.node.Right}
	right.postorderTraversal()
	node.node.PrintValue()
}

func main() {
	root := tree.Node{Value: 3}
	root.PrintValue()
	root.SetValue(2)
	root.PrintValue()

	root.Left = &tree.Node{Value: 1}
	root.Right = &tree.Node{Value: 3}
	fmt.Println("InorderTraverse")
	root.InorderTraverse()

	extendNode := myTreeNode{&root}
	fmt.Println("postorderTraversal")
	extendNode.postorderTraversal()

	// 值为nil也可以方法调用
	//var nilNode *Node
	//nilNode.SetValue(2)


}
