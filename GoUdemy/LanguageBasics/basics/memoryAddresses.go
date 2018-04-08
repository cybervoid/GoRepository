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
