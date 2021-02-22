package main

import "fmt"

// 
func consts(){
  const file = "abe.txt"
  const number1,number2 = 1,2
  fmt.Printf("number1: %s file: %s \n", number1, file)
}

// enums by iota
func enums(){
  const (
    cpp = iota
    java
    python
    golang
    javascript
  )

  const (
    b = 1 << (10 * iota)
    kb
    mb
    gb
    tb
    pb
  )

  fmt.Printf("%d,%d,%d,%d,%d\n", cpp,java,python,golang,javascript)
  fmt.Printf("%d, kb = %d , mb = %d, gb = %d, tb = %d, pb = %d \n", b,kb,mb,gb,tb,pb)
}


func main(){
  consts()
  enums()
}