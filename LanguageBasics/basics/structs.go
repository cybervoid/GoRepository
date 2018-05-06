package basics

import "fmt"

type person struct {
  first string
  last string
  age int
}

func StructsExample() {
  p1 := person{"James", "Bond", 20}
  p2 := person{"Jason", "Bourne", 31}
  fmt.Println(p1.first, p1.last, p1.age)
  fmt.Println(p2.first, p2.last, p2.age)
}
