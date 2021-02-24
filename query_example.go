package main

import "main/queue"

func main(){
  numbers := queue.Queue{1,2,3,4}
  numbers.Push(5)
  numbers.Print()

  numbers.Pop()
  numbers.Print()
}