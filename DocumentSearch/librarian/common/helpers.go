package common

import (
    "fmt"
    "log"
)

func Log(msg string) {
    log.Println("INFO - ", msg)
}

func Warn(msg string) {
    log.Println("------- helper.Warn() ---------")
    log.Println(fmt.Sprintf("Warn: %s", msg))
    log.Println("-------------------------------")
}
