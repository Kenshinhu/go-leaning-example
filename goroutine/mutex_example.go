package main

import (
  "fmt"
  "time"
  "sync"
)

func main() {
  
  var lock sync.Mutex
  
  go func(){
      
      lock.Lock()
      defer lock.Unlock()
      
      fmt.Printf("func1 get Locked @ %s \n", time.Now().String())
      
      time.Sleep(time.Second)
      
      fmt.Printf("func1 Unlock @ %s \n", time.Now().String())
    
  }()
  
  time.Sleep(time.Second)
  
  go func(){
    
    lock.Lock()
    defer lock.Unlock()
      
      fmt.Printf("func2 get Locked @ %s \n", time.Now().String())
      
      time.Sleep(time.Second)
      
      fmt.Printf("func2 Unlock @ %s \n", time.Now().String())
      
  }()
  
  
  time.Sleep(time.Second * 4)
  
  fmt.Println("Process is Exited!")
}