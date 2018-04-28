package basics

import (
  "time"
  "fmt"
)

func ChannelsBasicUnbufferedExample() {
  c := make(chan int)
  // c := make(chan int, 10) <-- buffered channel example
  go func() {
    for i := 0; i < 10; i++ {
        c <- i
    }
  }()

  go func() {
    for {
      x := <- c
      fmt.Println(x)
      if(x > 9) {
        break
      }
    }
  }()
  time.Sleep(time.Second)
}
