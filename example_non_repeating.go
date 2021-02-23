package main

import (
  "fmt"
  "unicode/utf8"
)

// Longest Substring Without Repeating Characters 
func notRepaetingSubString(s string) int{
  lastOccurred := make(map[byte]int)
  start := 0
  maxLength := 0; 

  for index, value := range []byte(s) {
    if count, exists := lastOccurred[value]; exists && count >= start {
      start =  count + 1
    }

    if index - start + 1 > maxLength {
      maxLength = index - start + 1;
    }
    lastOccurred[value] = index
  }
  return maxLength
}
 


func main(){
  fmt.Println("hello world")
  fmt.Println("abcABC=", notRepaetingSubString("abcAbc"))
  fmt.Println("abbcc=", notRepaetingSubString("abbcc"))
  
  fmt.Println(utf8.RuneCount([]byte("你好")))

  for i, ch := range []rune("你好") {
    fmt.Printf("(%d, %c)",i, ch)
  }


  fmt.Println(utf8.DecodeLastRuneInString(" "))
  fmt.Println(len("你好"))
}