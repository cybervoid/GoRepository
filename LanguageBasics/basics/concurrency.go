package basics

import (
  "fmt"
  "sync"
  "time"
  "math/rand"
  "sync/atomic"
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

var atomicityCounter int64 //int64 is a sign of atomicity
func AtomicityExample(){
  waitGroup.Add(2)
  go atomicityIncrementor("Foo: ")
  go atomicityIncrementor("Bar: ")
  waitGroup.Wait();
  fmt.Println("Counter Final: ", atomicityCounter)
}

func atomicityIncrementor(s string) {
  for i := 0; i < 20; i++ {
    time.Sleep(time.Duration(rand.Intn(3)) *time.Millisecond)
    atomic.AddInt64(&atomicityCounter, 1);
    fmt.Println(s, i, "Counter: ", atomicityCounter)
  }
  waitGroup.Done()
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
