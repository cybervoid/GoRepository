package basics

import (
  "fmt"
  "sync"
  "time"
  "math/rand"
)

var waitGroup sync.WaitGroup
var mutex sync.Mutex
var sharedCounter int
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
  go mutexIncrementor("Foo: ")
  go mutexIncrementor("Bar: ")
  waitGroup.Wait()
}

//helper methods
func mutexIncrementor(s string) {
  for i := 0; i < 20; i++ {
    time.Sleep(time.Duration(rand.Intn(20)) *time.Millisecond)
    mutex.Lock()
    x := sharedCounter
    x++
    sharedCounter = x
    fmt.Println(s, i, "Counter: ", sharedCounter)
    mutex.Unlock()
  }
  waitGroup.Done()
}

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
