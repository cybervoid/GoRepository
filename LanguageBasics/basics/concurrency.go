package basics

import (
  "fmt"
  "sync"
  "time"
)

var waitGroup sync.WaitGroup
var mutex sync.Mutex
//Concepts
func RunConcurrency() {
  runWithConcurrency()
}

func RaceCondition() {
  waitGroup.Add(2)
  go incrementor("Foo: ")
  go incrementor("Bar: ")
  waitGroup.Wait()
  fmt.Println("Final Counter: ", counter)
}

func MutexExample() {
  waitGroup.Add(2)
  go incrementor("Foo: ")
  go incrementor("Bar: ")
}

//helper methods
func runWithConcurrency() {
  waitGroup.Add(2)
  go foo()
  go bar()
  waitGroup.Wait()
}

func foo() {
  for i := 0; i < 10; i++ {
    fmt.Println("Foo: ", i)
    time.Sleep(time.Duration(3 * time.Millisecond))
  }
  waitGroup.Done()
}

func bar() {
  for i := 0; i < 10; i++ {
    fmt.Println("Bar: ", i)
    time.Sleep(time.Duration(4 * time.Millisecond))
  }
  waitGroup.Done()
}
