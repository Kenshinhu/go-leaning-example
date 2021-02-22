package main

import "fmt"

func ArrayPrinter(array [5]int){
  array[0] = 100
  for i, v := range array {
    fmt.Printf("index:%d value:%d \n", i , v)
  }
}

// Specify index replacement elements
func setItem(array *[5]int, target int, newItem int){
  array[target] = newItem;
}


func main(){
  var array1 [5]int
  array2 := [5]int{}
  array3 := [5]int{1,2,3,4,5}

  var grid [4][5]bool // 

  fmt.Println(array1)
  fmt.Println(array2)
  fmt.Println(array3)
  fmt.Println(grid)


  fmt.Println("Iterator Array Item")
  fmt.Printf("array3 : len = %d \n", len(array3))
  ArrayPrinter(array3)
  fmt.Printf("array3 = %v \n", array3)
  
  setItem(&array3, 1, 100)
  fmt.Printf("set 1 index to  100 : %v", array3)

}