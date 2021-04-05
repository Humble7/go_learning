package tree

func (node *Node) InorderTraverse() {
	if node == nil {
		return
	}

	node.Left.InorderTraverse()
	node.PrintValue()
	node.Right.InorderTraverse()
}
