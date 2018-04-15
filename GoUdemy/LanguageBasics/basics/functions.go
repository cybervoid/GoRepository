package basics

import "fmt"

func Greet(name string) {
  fmt.Println("Hello " + name)
}
func GreetFull(firstname string, lastname string) {
  fmt.Println("Hello " + firstname + " " + lastname)
}

func ReturnExample(name string) string {
  return fmt.Sprint(name)
}

func ReturnAssigned(fname string, lname string) (full string) {
  full = fmt.Sprint(fname, lname);
  return
}

func ReturnMultiple(fname, lname string) (string, string) {
  return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}

func FuncAverageVariadicParams(sf ...float64) float64 {
  total := 0.0
  for _, v := range sf {
    total += v
  }
  return total / float64(len(sf))
}


func FuncAverageVariadicArgs() {
  data := []float64 { 43, 44, 53, 64, 12 }
  n := FuncAverageVariadicParams(data...)
  fmt.Println(n)
}
