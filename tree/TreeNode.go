package tree

import "fmt"

type TreeNode struct {
    Value       int
    Left, Right *TreeNode
}

func (node *TreeNode) setValue(value int){
  node.Value = value
  node.Print()
}

// definition the function 
func (node *TreeNode) Print() {
  if node == nil {
    return
  }

  fmt.Println(node.Value)
}

func (node *TreeNode) Travel(){
  if node == nil {
    return
  }

  node.Left.Travel()
  node.Print()
  node.Right.Travel()
}

// factory function 
func CreateNode(value int) *TreeNode {
  return &TreeNode{Value: value}
}
