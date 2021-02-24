package tree

import "fmt"

type TreeNode struct {
    Value       int
    Left, Right *TreeNode
}

func (node *TreeNode) setValue(value int){
  node.Value = value
  node.print()
}

// definition the function 
func (node TreeNode) print() {
  fmt.Println(node.Value)
}

func (node *TreeNode) Travel(){
  if node == nil {
    return
  }

  node.Left.Travel()
  node.print()
  node.Right.Travel()
}

// factory function 
func CreateNode(value int) *TreeNode {
  return &TreeNode{Value: value}
}
