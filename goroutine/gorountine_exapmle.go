package main

import (
  "fmt"
  "time"
)



func setValue(v *int, value int){
  *v = value
}

func main(){
  
  v := new(int)
  
  
  go setValue(v,1)
  go setValue(v,2)
  
  for i :=0 ; i< 10 ; i++ {
    time.Sleep(time.Second)
    fmt.Printf("v=%d \n",*v)
  }
  // Depending on the scheduling of the debugger, the result can be 0, 1, 2
  
}