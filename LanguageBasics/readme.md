
## Go commands
1. `go run` - run a file ad hoc
2. `go install` - installs the go program
3. `go clean ./bin/` - clean's extra "stuff" in directories with leftover build content
4. `go build /path/main.go`   - builds the go program

## Language Syntax
1. Public and Private - Capitalize variables and methods to make them public. Lowercase for private.
2. [Variables](./basics/vars.go) - also see example 2 below
3. Importing packages
4. [Blank Identifier](./basics/blank.go) - Also see Example 4 below
5. Constants - see example 5 below
6. [Memory addresses](./basics/memoryAddresses.go) - Also see Example 7 below
7. [Pointers](./basics/pointers.go) - Also see Example 7 below
8. [Scope](./basics/scope.go)
9. [Control Flow](./basics/controlflow.go) - Includes If statements, loops, switch statements
10. [Runes](./basics/runes.go)
11. [Strings](./basics/stringthings.go)
12. [Functions](./basics/functions.go) - Includes callbacks, expressions, anonymous expressions, parameters, variadic arguments, variadic params.
  1. [Callback Example #2](./examples/filtercallback.go)

Example 2 - Variables
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
- slices

```
xs := []int{} // declare an empty slice
xs = append(xs, 5) // add int 5 to slice xs  
```

Example 3 - Imports
```
import (
  "fmt"
  "somepackage"
)
or;

import "fmt"
```
Example 4 - Blank Identifier
```
res, _ := http.Get("https://blockexplorer.com/api/addr/1Q85UwrAUKm4EXE5jSmpoyTiys8BCos45J")
```
Example 5 - Constants
```
const p string = "my constant example"
```
Example 6 - Memory addresses
```
a := 43
fmt.Println("a - ", a)
fmt.Println("a's memory address - ", &a)
```
Example 7 - Pointers
```
a := 43
fmt.Println(a)
fmt.Println(&a)
//b is a pointer, b points to a's memory address
var b *int = &a
fmt.Println(b) //print the memory address
fmt.Println(*b) //print the value at the memory address with the * operator
```
