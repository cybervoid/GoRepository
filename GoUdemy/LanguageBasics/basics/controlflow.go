package basics

import "fmt"

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
