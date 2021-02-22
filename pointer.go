package main

import "fmt"

func swap( a, b *int){
  *a, *b = *b, *a
}

func main() {
  a := 1
  b := 2 
  fmt.Printf(" a = %d, b = %d \n", a,b)
  swap(&a, &b)
  fmt.Printf("swapping \n a = %d, b = %d", a, b )
}