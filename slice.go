package main

import "fmt"

func main() {
  
  var slice1 []int
  slice2 := []int{1,2,3,4,5}
  slice := make()
  slice1 = slice2

  fmt.Println(slice1)

}