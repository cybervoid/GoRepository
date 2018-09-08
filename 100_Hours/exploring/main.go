package main

import (
    "fmt"
    "time"
)

//var waitGroup sync.Waitgroup

func main() {
    myChan := make(chan int, 10)
    go func() {
        for i := 0; i < 100; i++ {
            myChan <- i
        }
    }()

    go func() {
        for i := 100; i > 0; i-- {
            myChan <- i;
        }
        myChan <- 101;
    }()

    go func() {
        for {
            x := <- myChan
            fmt.Println(x)
            if(x > 100) {
                break
            }
        }
    }()
    time.Sleep(time.Second)
}
