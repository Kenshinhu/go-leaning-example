package main

import (
  "fmt"
  "os"
  "bufio"
)


func printInput(ch chan string){
  for input := range ch {
    if input == "EOF" {
      break
    }
    fmt.Printf("input is : %s \n",input)
  }
}

func main(){
  
  ch := make(chan string)
  defer close(ch)
  
  go printInput(ch)
  
  // read io from input 
  scanner := bufio.NewScanner(os.Stdin)
  
  for scanner.Scan() {
    
    val := scanner.Text()
    
    ch <- val
    
    if val == "EOF" {
      fmt.Printf("End the game! \n")
      break
    }
    
  }
  
}