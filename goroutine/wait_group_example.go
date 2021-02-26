package main

import (
  "fmt"
  "time"
  "strconv"
  "sync"
)

func main(){
  
  var waitGroup sync.WaitGroup
  
  groupLens := 5
  
  waitGroup.Add(groupLens)
  
  for i:=0; i< groupLens; i++ {
    go func(index int){
      fmt.Println("work "+strconv.Itoa(index)+" is done @" + time.Now().String())
      
      time.Sleep(time.Second * 1)
      
      waitGroup.Done()
    }(i)
  }
  
  waitGroup.Wait()
  
  fmt.Println("all workd are done @"+time.Now().String())
}