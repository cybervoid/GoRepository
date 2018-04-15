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
