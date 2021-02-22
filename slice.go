package main

import "fmt"

func updateSlice(array []int, target int, newItem int) []int {
  array[target] = newItem
  return array
}

func removeItem(array []int, target int) []int{
  return append(array[:target], array[target+1:]...)
}

func slicePrinter(slice []int){
  fmt.Printf("values=%v , len = %d , cap = %d ", slice, len(slice), cap(slice))
}

func pop(slice []int) []int{
  return slice[:len(slice) - 2]
}



func main() { 

  // Variable assignment
  arr := [...]int{0,1,2,3,4,5}

  fmt.Printf("%v \n",arr)

  // get the slice from 2 to 5
  fmt.Printf("Array[2:5] = %v \n", arr[2:5])

  // get the slice from 0 to 5
  fmt.Printf("Array[:5] = %v \n", arr[:5])

  // get the slice 
  fmt.Printf("Array[:] = %v \n", arr[:])

  fmt.Printf("After Update Slice the 100 at index = 1 \n")
  updateSlice(arr[:], 1, 100)
  fmt.Printf("Array[:] = %v \n", arr[:])

  array1 := arr[1:5]
  fmt.Printf("array1[1:5] = %v \n", array1)

  // slice 2 get item for Out of Bounds 
  array2 := array1[2:4]
  fmt.Printf("array2[:5] = %v arrays's, len = %d, cap = %d \n", array2, len(array2), cap(array2))


  // append element 
  array3 := append(array2, 6)
  fmt.Printf("array3 = %v arrays's, len = %d, cap = %d \n", array3, len(array3), cap(array3))


  // create slice from make function
  array4 := make([]int, 10 ,12) // create the slice len= 10 , cap = 12
  array4 = append(array2[:1],array3[0:]...)
  fmt.Printf("array4 = %v arrays's, len = %d, cap = %d \n", array4, len(array4), cap(array4))

  fmt.Printf("array4 remove index 2 %v \n",removeItem(array4, 2))

  fmt.Printf("pop array4 : %v",pop(array4))
}