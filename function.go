package main

import (
  "fmt"
  "runtime"
  "reflect"
)

func div(a,b int) (int, int) {
  return a / b, a % b;
}

func apply(op func(int,int) (int,int), a,b int) (int,int) {
  p := reflect.ValueOf(op).Pointer() // reflect the function , get the function name
  opName := runtime.FuncForPC(p).Name()
  fmt.Printf("Calling functinon : %s with args(%d, %d) \n", opName, a,b)
  return op(a,b)
}




func eval(a,b int, op string) (int ,error) {
  switch op {
    case "+":
      return a+b, nil
    case "-": 
      return a-b, nil
    case "*":
      return a*b, nil
    case "/":
      result, _ := div(a, b)
      return result, nil
    default:
      return 0, fmt.Errorf("unsupprted operation: %s", op)
  }
}

func main() {
  result, _ := eval(7, 6, "/")
  fmt.Printf(" 7 / 6 = %d \n", result)

  q, r := apply(div,11,6)
  fmt.Printf(" apply: 11 / 6 = %d, %d \n ", q,r)


}