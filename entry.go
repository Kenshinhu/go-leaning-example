package main

import "main/tree"

func main(){

	var root tree.TreeNode
	left := tree.TreeNode{Value: 3}
	root.Left = &left
	root.Right = tree.CreateNode(45)

	root.Travel()
}