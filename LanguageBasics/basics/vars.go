package basics

import "fmt"


func ShortHand() {
  fmt.Println("variables as shorthand")
  a := 10
  b := "golang"
  c := 4.17
  d := true

  fmt.Printf("%v \n", a)
  fmt.Printf("%v \n", b)
  fmt.Printf("%v \n", c)
  fmt.Printf("%v \n", d)
}

func AsVars() {
  fmt.Println("As statically typed variables")
  var aa int32  = 12
  var bb, b2 string = "golang_2", "stored in b2"
  var cc float32 = 4.172
  var dd bool = true
  fmt.Printf("%v \n", aa)
  fmt.Printf("%T \n", bb)
  fmt.Printf("%v \n", b2)
  fmt.Printf("%v \n", cc)
  fmt.Printf("%T \n", dd)
}
