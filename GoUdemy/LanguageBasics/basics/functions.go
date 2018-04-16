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


func ExpressionExample(){
  //anonymous function example
  //when put into a variable is called an expression
  greeting := func() {
    fmt.Println("Hello World")
  } // must always be put into a variable
  greeting()
  fmt.Printf("%T\n", greeting)
}

func makeGreeter() func() string {
  return func() string {
    return "hello world"
  }
}

func ClosureExpression() {
  greets := makeGreeter()
  fmt.Println(greets())
  fmt.Printf("%T\n", greets)
}

func incrementer(x int) int {
  x++
  return x
}

func ClosureFunc(){
  var x int = 0
  increment := func() int {
    x++
    return x
  }
  fmt.Println(increment())
  fmt.Println(increment())
}

func wrapper() func() int {
  x := 0
  return func() int {
    x++
    return x
  }
}

func ClosureNotAnon() {
  increment := wrapper()
  fmt.Println(increment())
  fmt.Println(increment())
}

//callback: passing a func as an argument
func callbacks(numbers []int, callback func(int)) {
  for _, n := range numbers {
    callback(n)
  }
}
func RunCallbacks(){
  callbacks([]int{ 10, 3, 43, 44, 76, 7 }, func(n int){
    fmt.Println(n)
  })
}
