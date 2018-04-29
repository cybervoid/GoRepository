package basics

import "fmt"

type Contact struct{
  Greeting string
  Name string
}

func ForLoop() {
  sum := 0
  for i := 0; i < 10; i++ {
    sum += i
  }
  fmt.Println(sum)
}

func WhileLoop(){

  var cond = true
  var counter = 0
  for cond {
    counter++
    fmt.Println(counter)
    if counter > 10{
      cond = false
    }
  }
}

func BreaksConditions(){
  var cond = true
  var counter = 0
  for cond {
    counter++
    fmt.Println(counter)
    if counter > 10{
      break
    }
  }
}



func SwitchStatements(){
  switch "Ryan" {
  case "NotRyan":
      fmt.Println("Hello not ryan")
    case "Ryan":
      fmt.Println("Hello Ryan")
    default:
      fmt.Println("Default!")
  }
}
func SwitchFallthroughStatements(){
  switch "Ryan" {
  case "NotRyan":
      fmt.Println("Hello not ryan")
    case "Ryan":
      fmt.Println("Hello Ryan")
      fallthrough
    case "Fallthrough":
      fmt.Println("fallthrough executed!")
    default:
      fmt.Println("Default!")
  }
}


func SwitchOnType(x interface{}){
  switch x. (type) { //this is an assert; asserting x is of type
  case int:
    fmt.Println("int")
  case string:
    fmt.Println("string")
  case Contact:
    fmt.Println("contact")
  default:
    fmt.Println("unknown")
  }
}

func IfStatements(cond bool){
  if cond == true {
    fmt.Println("it was true!")
  } else {
    fmt.Println("it was false!")
  }
}
func IfInitializationStatements(cond bool){
  b := true
  if food := "Chocolate"; b {
    fmt.Println(food)
  } else {
    fmt.Println("not food")
  }
}
