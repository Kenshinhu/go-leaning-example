package main

import (
  "fmt"
  "time"
)

func send(sender chan int, begin int, limit int){
  for i :=begin; i< limit; i++ {
   sender <- i 
  }
}

func main(){
  
  ch1 := make(chan int)
  ch2 := make(chan int)
  
  go send(ch1, 0, 10)
  go send(ch2, 10, 20)
 
  time.Sleep(time.Second) 
  
  for {
    select {
      case val := <- ch1 :
        fmt.Printf("Get Value  %d from ch1 \n", val)
      case val := <- ch2 :
        fmt.Printf("Get Value  %d from ch2 \n", val)
      case <- time.After(2 * time.Second):
        fmt.Println("Time Out!!")
        return;
    }
  }
  
  
}