package basics

import (
  "runtime"
  "time"
  "fmt"
  "math/rand"
)

var counter int
// this line is not required after GO 1.5
func init() { //init is a special function that runs before the rest of the code
  runtime.GOMAXPROCS(runtime.NumCPU())
}
func RunParallel() {
  waitGroup.Add(2)
  go foo()
  go bar()
  waitGroup.Wait()
}


func RaceCondition() {
  waitGroup.Add(2)
  go incrementor("Foo: ")
  go incrementor("Bar: ")
  waitGroup.Wait()
  fmt.Println("Final Counter: ", counter)
}

func incrementor(s string) {
  for i := 0; i < 20; i++ {
    x := counter
    x++
    time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
    counter = x
    fmt.Println(s, i, "Counter: ", counter)
  }
  waitGroup.Done()
}
