package main

import "main/tree"

type myTreeNode struct {
  node *tree.TreeNode
}

func (myNode *myTreeNode) postOrder(){
  if myNode.node == nil {
    return 
  }

  left := myTreeNode{myNode.node.Left};
  left.postOrder()
  right := myTreeNode{myNode.node.Right}
  right.postOrder()
  myNode.node.Print()
}



func main(){
  var root tree.TreeNode
  left := tree.TreeNode{Value: 3}
  root.Left = &left
  root.Right = tree.CreateNode(45)
  root.Travel()

  mytree := myTreeNode{&root}
  mytree.postOrder()
}