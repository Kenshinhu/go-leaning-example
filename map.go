package main

import "fmt"

func main(){   
  //  set map with default value
  m := map[string]string{
    "john": "english",
    "amy": "chinese",
    "kin": "mathematics",
    "ouban":"physical",
  }

  fmt.Println(m)

  // variable the empty map
  m2 := make(map[string]string) // empty map
  fmt.Println(m2)

  var m3 map[string]string // m3 = nil
  fmt.Println(m3)

  // travering maps
  for k,v := range m {
    fmt.Println(k,"\t",v)
  }

  // Get the value of the specified key
  fmt.Println("john","=",m["john"])

  // check the key is exists
  if _, exists := m["decade"]; exists == false {
    fmt.Println("key 'decade' is not exists")
  }


  // Set the value 
  m["kitty"] = "chemical"
  fmt.Println("kitty","=",m["kitty"])

  // delete item by key
  delete(m, "kitty")
  _, kittyExists := m["kitty"]
  if kittyExists == false {
    fmt.Println("kitty is not exists! ")
  }

  

}