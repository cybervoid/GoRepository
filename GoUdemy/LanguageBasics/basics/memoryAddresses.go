package basics

import "fmt"

func MemAddress(){
  a := 43
  fmt.Println("a - ", a)
  fmt.Println("a's memory address - ", &a)
}
const metersToYards float64 = 1.09361
func UseMemAddress(){
  var meters float64
  fmt.Print("Enter meters swam: ")
  fmt.Scan(&meters)
  yards := meters * metersToYards
  fmt.Println(meters, " meter is ", yards, " yards.")
}

func UseMemAddressAgain(){
  x := 5
  fmt.Println(x)
  fmt.Printf("%p\n", &x) //address in func UseMemAddressAgain
  fmt.Println(&x) //address in func UseMemAddressAgain
  zero(&x)
  fmt.Println(x)
}

func zero(x *int){
  fmt.Printf("%p\n", &x) //address in func zero
  fmt.Println(&x) //address in func zero
  *x = 0
}
