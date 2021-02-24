package main

import "fmt"

type treeNode struct {
    value int
    left, right *treeNode
}

func (node *treeNode) setValue(value int){
  node.value = value
  node.print()
}

// definition the function 
func (node treeNode) print() {
  fmt.Println(node.value)
}

func (node *treeNode) travel(){
  if node == nil {
    return
  }

  node.left.travel()
  node.print()
  node.right.travel()
}

// factory function 
func createNode(value int) *treeNode {
  return &treeNode{value: value}
}

func main(){

  // 3 type for definition object
  var root treeNode
  left := treeNode{value:3}
  root.right = &treeNode{value:5}
  root.left = &left
  root.right.left = new(treeNode)
  root.right.left = createNode(10)

  // defintion slice the objects
  nodes := []treeNode{
    {value: 1},
    {},
    {value: 2},
  }
  
  fmt.Println(root)

  fmt.Println(nodes)
 
  root.travel()
}