package main
import (
  "fmt" 
  "time"
)

// 'chan <- int' means that only data can be sent to the channel
// '<- chan int' means that only data can be revice from the channel
// 'chan int' means that Two-way data flow
func worker(i int, c chan int) {
  for n := range c {
    fmt.Printf("Worked %d received : %c\n",i, n)
  }
} 

func createWork(id int) chan <- int {
  c := make(chan int)
  go worker(id, c)
  return c
}


func main(){
  var channel[10] chan <- int

  for i:=0 ; i<10; i++ {
    channel[i] = createWork(i)      
    channel[i] <- 'a' + i
    
  }  

  time.Sleep(time.Millisecond)

}