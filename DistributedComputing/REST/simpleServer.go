package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
)
var counter int = 0;
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
    msg := fmt.Sprintf("Received request [%s] for path: [%s]", r.Method, r.URL.Path)
    log.Println(msg)
    var s string = strconv.Itoa(counter)
    response := fmt.Sprintf("Hello, world!  at path: [%s],  counter: [%v]", r.URL.Path, s)
    fmt.Fprintf(w, response)
    counter++
}

func main() {
    http.HandleFunc("/", helloWorldHandler) //catch all Path
    log.Println("starting server at port :8081")
    http.ListenAndServe(":8081", nil)
}
