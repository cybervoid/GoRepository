package basics

import "fmt"

func PointerExample(){
  a := 43
  fmt.Println(a)
  fmt.Println(&a)
  //b is a pointer, b points to a's memory address
  var b *int = &a
  fmt.Println(b)

}
