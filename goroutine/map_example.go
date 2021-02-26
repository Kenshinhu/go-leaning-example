package main

import (
  "fmt"
  "sync"
  "strconv"
  )

var syncMap sync.Map
var waitGroup sync.WaitGroup

func addNumber(start int)  {
  // start = 0; 0,1,2
  // start = 10; 10,11,12
  // start = 20; 20,21,22
  // start = 30; 30,31,32
  // start = 40; 40,41,42
  for i:=start; i< start + 3 ; i++ {
    syncMap.Store(i,i)
  }
  
  waitGroup.Done()
  
}
  

func main(){
  routineSize := 5
  
  waitGroup.Add(routineSize)
  
  for i := 0 ; i<routineSize ; i++ {
    go addNumber(i * 10)
    
  }
  
  waitGroup.Wait()
  
  var size int
  
  syncMap.Range(func (key, value interface{}) bool {
    
    size ++ 
    
    return true
    
  })
  
  fmt.Println("syncMap current size is : "+ strconv.Itoa(size))
  
   value, ok := syncMap.Load(0); 
   
   if ok {
    fmt.Println("key 0 has Value", value)
  }
  
  
  
  
  fmt.Println("map_example.go")
}