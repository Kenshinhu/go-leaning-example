package main
import "fmt"
// definition Gobal variable
var (
  number1 = 1
  string1 = "2"
  bool1 = true
)

func variable_type_1() {
  var number int
  var str string
  fmt.Printf("number:%d string:%s", number, str)
} 

// initial value 
func variableInitialValue(){
  var number1, number2 int = 3,4
  var str1, str2 string = "string1", "string2"
  fmt.Printf("number1 : %d, number2: %d, string1: %s, string2: %s \n",number1, number2, str1, str2)
}

// initial value by type deduction
func variableInitalValueByTypeDeduction(){
  var number,str = 1, "hello world"
  fmt.Printf("number:%d string:%s \n", number, str)
}

// initial value by shorter
func variableInitialShorter(){
  number1, string1, bool1 := 1,"2",true
  fmt.Printf("number:%d string:%s bool:%b \n", number1, string1, bool1)
}

func main() { 
  variable_type_1()
  variableInitialValue()
  variableInitalValueByTypeDeduction()
  variableInitialShorter()

  fmt.Printf("Global = number1:%d string1:%s bool1:%b", number1, string1, bool1)
}