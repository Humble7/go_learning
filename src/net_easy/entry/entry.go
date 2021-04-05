package main

import (
	"fmt"
	"net_easy/tree"
)

func main() {
	root := tree.Node{Value: 3}
	root.PrintValue()
	root.SetValue(2)
	root.PrintValue()

	root.Left = &tree.Node{Value: 1}
	root.Right = &tree.Node{Value: 3}
	fmt.Println("InorderTraverse")
	root.InorderTraverse()

	// 值为nil也可以方法调用
	//var nilNode *Node
	//nilNode.SetValue(2)


}
