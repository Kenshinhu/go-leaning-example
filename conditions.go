package main

import (
  "fmt"
  "io/ioutil"
)

// switch coding
func eval(a,b int, op string) int{
  var results int
  switch op {
    case "+":
      results = a + b
    case "-":
      results = a - b
    case "*":
      results = a * b
    case "/":
      results = a / b
    default:
      panic("unsupported operator: "+ op)
  }
  return results
}

func grade(score int) string {
    g := ""
    switch{
      case score < 0 && score > 100:
        panic(" Wrong score")
      case score < 60:
        g = "F"
      case score < 80:
        g = "C"
      case score < 90:
        g = "B"
      case score < 100:
        g = "A"
    }
    return g
}


func readFile(){
  const filename = "abc.txt"
  if contents, err := ioutil.ReadFile(filename) ; err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("contents : %s \n", string(contents))
  }
}

func main(){
  readFile()
  fmt.Printf(" 4 + 5 = %d \n",eval(4,5,"+"))
  fmt.Printf(" 2 - 1 = %d \n",eval(2,1,"-"))
  fmt.Printf(" 2 * 2 = %d \n",eval(2,2,"*"))
  fmt.Printf(" 10 / 2 = %d \n",eval(10,2,"/"))

  fmt.Printf("John grade: %s \n",grade(50))
}