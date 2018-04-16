package basics

import (
  "runtime"
)
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
