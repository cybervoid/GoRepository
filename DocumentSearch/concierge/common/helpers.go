package common

import (
    "fmt"
    "log"
    "regexp"
    "strings"
)

// Log is used for simple logging to console
func Log(msg string) {
    log.Println("Info - ", msg)
}

// warn is used to to log warning messages to console.
func Warn(msg string) {
    log.Println("--------------------")
    log.Println(fmt.Sprintf("WARN: %s", msg))
    log.Println("--------------------")
}

var punctuations = regexp.MustCompile('^\p{P}+|\p{P}+$')
