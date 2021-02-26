package main

import (
  "fmt"
  "time"
)

func consume(ch chan int){
  time.Sleep(time.Second * 10) // sleep 1 second
  
  <- ch // cosume the channel data
}

func main(){
  
  n := make(chan int, 2) // make the 2 length to channel
  
  go consume(n)
  
  n <- 1
  n <- 2
  fmt.Println("I'm Free ")
  
  n <- 3
  fmt.Println("Hold on")
  
  time.Sleep(time.Second)
  
}