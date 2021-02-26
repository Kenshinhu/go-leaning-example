package main

import (
  "fmt"
  "time"
  "strconv"
  "sync"
)

var rwlock sync.RWMutex

func main(){
  for i:=0; i< 5 ; i++ {
    go func(i int){
      rwlock.RLock()
      defer rwlock.RUnlock()
      
      fmt.Println("Read func "+strconv.Itoa(i)+" get lock @"+time.Now().String())
      
      time.Sleep(time.Second)
    }(i)
  }
  
  time.Sleep(time.Second)
  
  
  // write Locker
  for i:=0; i< 5 ; i++ {
    go func(i int){
      rwlock.Lock()
      defer rwlock.Unlock()
      
      fmt.Println("Write func "+strconv.Itoa(i)+" get lock @"+time.Now().String())
      
      time.Sleep(time.Second)
    }(i)
  }
  
  time.Sleep(time.Second)
  
}