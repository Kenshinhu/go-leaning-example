package main

import (
  "fmt"
  "time"
  "runtime"
)

func main(){
  var a [10]int
  for i := 0; i< 10; i++ {
    go func(i int){
      for {
          a[i]++
          runtime.Gosched() // 
      }
    }(i)
  }

  time.Sleep(time.Millisecond)

  fmt.Println(a)

  // fmt.Println("Hello World")
}