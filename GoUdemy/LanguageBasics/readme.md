
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
- as variables var
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
