package more

import (
  "fmt";
)

func CountEvenOdd() {
  var (
    even int = 0
    odd int = 0
  )
  numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
  for _; n := range numbers {
    if n % 2 == 0{
      even += 1
    } else {
      odd += 1
    }
  }
  fmt.Printf("even %d\n odd %d\n", even, odd)
}

func MakeLabel() {
  MyLabel:
    for {
      break MyLabel
    }
}
