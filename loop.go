package main

import (
  "fmt"
  "strconv"
  "bufio"
  "os"
)

func covertToBin(n int) string{
    result := ""
    for ; n > 0; n/=2 {
      lsb := n % 2
      result = strconv.Itoa(lsb) + result
    }
    return result
}

func readLine(filename string){
  file, err := os.Open(filename)
  if err != nil {
    panic(err)
  }

  scanner := bufio.NewScanner(file)

  for scanner.Scan()  {
    fmt.Println(scanner.Text())
  }
}

func loopEvent(){
  for {
    fmt.Println("abc")
  }
}

func main(){
  fmt.Println(covertToBin(13));
  fmt.Println(covertToBin(234));
  fmt.Println(covertToBin(234123));
  readLine("text.txt")
  loopEvent();
}