
## Go commands
1. `go run` - run a file ad hoc
2. `go install` - installs the go program
3. `go clean ./bin/` - clean's extra "stuff" in directories with leftover build content
4. `go build /path/main.go`   - builds the go program
#### In pkg folder
5. Note system architecture, same commands can run

## Language Syntax
1. Capitalize variables and methods to make them public
2. Declaring variables:
 - shorthand (only inside of a func):
```
Implicit type inference
a := 10
b := "golang"
c := 4.17
d := true

fmt.Printf("%v \n", a)
fmt.Printf("%v \n", b)
fmt.Printf("%v \n", c)
fmt.Printf("%v \n", d)
```
- var
```
var ()


defaults: 0, "", nil, false  (nil not null)
```
3. Importing packages
```import (
  "fmt"
)```
