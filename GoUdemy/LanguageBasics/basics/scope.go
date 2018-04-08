package basics
import "fmt"

var packageLevelScopeExample = 44

func scope(){
  var funcLevelScopeExample = "example" //also a block scope area, aka closure
  fmt.Println(funcLevelScopeExample)
}

func Closure(){
  x := 42
  fmt.Println(x)
  //no relevance to previous line
  {
    fmt.Println(x)
    y :=  "quick brown fox "
    fmt.Println(y)
  }
}
