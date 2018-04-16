
## Go commands
1. `go run` - run a file ad hoc
2. `go install` - installs the go program
3. `go clean ./bin/` - clean's extra "stuff" in directories with leftover build content
4. `go build /path/main.go`   - builds the go program
#### In pkg folder
5. Note system architecture, same commands can run

## Language Syntax
1. Capitalize variables and methods to make them public
2. [Variables](./basics/vars.go)
```
var aa int32  = 12
var bb, b2 string = "golang_2", "stored in b2"
var cc float32 = 4.172
var dd = true

defaults: 0, "", nil, false  (nil not null)
```
- shorthand (only inside of a func):
```
Implicit type inference
a := 10
b := "golang"
c := 4.17
d := true
```
3. Importing packages
```
import (
  "fmt"
  "somepackage"
)
or;

import "fmt"
```
4. [Blank Identifier](./basics/blank.go)
```
res, _ := http.Get("https://blockexplorer.com/api/addr/1Q85UwrAUKm4EXE5jSmpoyTiys8BCos45J")
```
5. Constants
6. [Memory addresses](./basics/memoryAddresses.go)
```
a := 43
fmt.Println("a - ", a)
fmt.Println("a's memory address - ", &a)
```
7. [Pointers](./basics/pointers.go)
```
a := 43
fmt.Println(a)
fmt.Println(&a)
//b is a pointer, b points to a's memory address
var b *int = &a
fmt.Println(b) //print the memory address
fmt.Println(*b) //print the value at the memory address with the * operator
```
8. [Scope](./basics/scope.go)
9. [Control Flow](./basics/controlflow.go) - Includes If statements, loops, switch statements
10. [Runes](./basics/runes.go)
11. [Strings](./basics/stringthings.go)
12. [Functions](./basics/functions.go) - Includes callbacks, expressions, anonymous expressions, parameters, variadic arguments, variadic params.
