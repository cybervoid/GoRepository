package main

import (
	"flag"
	"fmt"
	"os"
    "encoding/csv"
    "strconv"
    "bufio"
)
//"strings"
//"encoding/csv"
//"bufio"
//"io"

type problem struct {
    problem string
    answer int32
}
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func fileExists(path string) bool {
    if _, err := os.Stat(path); err == nil {
        return true
    }
    return false
}

func loadFile() {
    var filepath = "quiz_problems.csv"
    csvFilename := flag.String("csv", filepath, "a csv file in the format of 'question, answer'")
    flag.Parse()
    file, err := os.Open(*csvFilename)
    check(err)
    r := csv.NewReader(file)
    lines, err := r.ReadAll()
    check(err)
    problems := parseLines(lines)
    for i, p := range problems {
        fmt.Printf("Problem #%d: %s = \n", i + 1, p.problem)

        buf := bufio.NewReader(os.Stdin('\n'))
        fmt.Print("> ")
    }
}

func parseLines(lines [][]string) []problem {
    ret := make([]problem, len(lines))
    for i, line := range lines {
        ret[i].problem = line[0]
        a, err := strconv.ParseInt(line[1], 10, 32)
        check(err)
        ret[i].answer = int32(a)
    }
    return ret
}

func main() {
    loadFile()
}
