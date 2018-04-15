package basics

import "fmt"

func StringLiterals(){
  name := "Ryan W"
  fmt.Println(`
    <!DOCTYPE html>
    <html lang="en">
    <head>
    <meta charset="UTF-8">
    <title>Go  Piggers!</title>
    </head>
    <body>
    <h1>` + name + `</h1>
    </body>
    </html>`)
}
