package main

import (
	"fmt"
)

type treeNode struct {
	value int
	left, right *treeNode
}

func main() {
	var root treeNode

	root = treeNode{value: 3}

	root.left = &treeNode{}

	root.right = &treeNode{5, nil, nil}

	root.right.left = new(treeNode)

	root.left.right = createTreeNode(2)

	//nodes := []treeNode {
	//	{value: 3},
	//	{},
	//	{6, nil, &root},
	//}



	//fmt.Println(nodes)

	//fmt.Println(root)

	root.print()
	root.right.left.setValue(4)
	root.right.left.print()

	//root.left.right = createTreeNode(2)

}

//工厂
func createTreeNode(value int) *treeNode {
	return &treeNode{value: value}
}

//接收
func (node treeNode) print() {
	fmt.Println(node.value)
}

func (node *treeNode) setValue(value int) {
	node.value = value
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}