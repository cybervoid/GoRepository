package basics

import (
  "time"
  "fmt"
  "sync"
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

func ChannelsUnbufferedCloseExample() {
  fmt.Println("start")
  c := make(chan int)
  // c := make(chan int, 10) <-- buffered channel example
  go func() {
    for i := 0; i < 10; i++ {
        c <- i
    }
    close(c)
  }()

  go func() {
    for n := range c {
      fmt.Println(n)
    }
  }()
  time.Sleep(time.Second)
}

func MultipleChannelsUnbufferedExample() {
  fmt.Println("start")
  c := make(chan int)
  var chanWg sync.WaitGroup
  chanWg.Add(2)
  go func() {
    //chanWg.Add(1)
    for i := 0; i < 10; i++ {
      c <- i;
    }
    chanWg.Done()
  }()

  go func() {
    //chanWg.Add(1)
    for i := 0; i < 10; i++ {
      c <- i
    }
    chanWg.Done()
  }()

  go func() {
    chanWg.Wait()
    close(c)
  }()

  for n := range c {
    fmt.Println(n)
  }
}

//A system of sending messages
func SemaphoreExample() {
  c := make(chan int)
  done := make(chan bool)

  go func() {
    for i := 0; i < 10; i++ {
      c <- i
    }
    done <- true
  }()

  go func() {
    for i := 0; i < 10; i++ {
      c <- i
    }
    done <- true
  }()

  go func() {
    <- done
    <- done
    close(c)
  }()
  for n := range c {
    fmt.Println(n)
  }
}
