package main

import (
  "../basics";
  "fmt"
)

func main() {
  //basics.ShortHand()
  //basics.AsVars()
  //basics.Closure()
  //basics.BlankIdentifierExample()
  //basics.MemAddress()
  //basics.UseMemAddress()
  //basics.PointerExample()
  //basics.UseMemAddressAgain()
  //basics.ForLoop()
  //basics.WhileLoop()
  //basics.Cast()
  //basics.RuneCounter()
  //basics.SwitchFallthroughStatements()
  //basics.SwitchOnType(5)
  //basics.SwitchOnType("testing")
  //var t = basics.Contact{ "Hello!", "Ryan!" }
  //basics.SwitchOnType(t)
  //basics.IfInitializationStatements(false)
  //basics.GreetFull("Bill", "Burr")
  //fmt.Println(basics.ReturnExample("this is Ryan"))
  //fmt.Println(basics.ReturnAssigned("Ryan ", "Rules"))
  //fmt.Println(basics.ReturnAssigned("Ryan ", "Rules"))
  //fmt.Println(basics.ReturnMultiple("Ryan", "Dubs"))
  //fmt.Println(basics.FuncAverageVariadicParams(5.554, 76.766, 123.674, 22.32223))
  //basics.FuncAverageVariadicArgs()
  //basics.ExpressionExample()
  //basics.ClosureExpression()
  //basics.ClosureFunc()
  //basics.ClosureNotAnon()
  //basics.RunCallbacks()
  //examples.RunFilter()
  //basics.RunRecursiveFactorial()
  //basics.RunDefer(false)
  //basics.RunPassByValue()
  //basics.RunPassByReference()
  //basics.RunPassByReferenceMap()
  //basics.RunAnonymousSelfExecuting()
  //basics.SliceExample()
  //basics.RunInterfacesExample()

  runExample(25)

  //printOptions()

  //fmt.Println(stringutil.MyName)
}
func runExample(selected int) {
  switch selected {
  case 18:
    fmt.Println("RunConcurrency()")
    basics.RunConcurrency()
  case 19:
    fmt.Println("RaceCondition()")
    basics.RaceCondition()
  case 20:
      fmt.Println("MapExample()")
      basics.MapExample()
  case 21:
    fmt.Println("MutexExample()")
    basics.MutexExample()
  case 22:
    fmt.Println("AtomicityExample()")
    basics.AtomicityExample()
  case 23:
    fmt.Println("ChannelsBasicExample()")
    basics.ChannelsBasicUnbufferedExample()
  case 24:
    fmt.Println("ChannelsUnbufferedCloseExample()")
    basics.ChannelsBasicUnbufferedExample()
  case 25:
    //fmt.Println("MultipleChannelsUnbufferedExample()")
    basics.MultipleChannelsUnbufferedExample()
  default:
    fmt.Println("Invalid option")
  }
}

// type option struct {
//   name string
//   function func
// }
//
// func createFuncOptions() (options map[int] *option) {
//   options = :make(map[int] *option)
//   var ex option := option{ name: "RaceCondition", function: basics.RaceCondition() }
//
//   options[1] = ex
//   return options
// }
//
// func printOptions() {
//   var opt := createFuncOptions()
//   fmt.Println(opt[1])
//
//   reader := bufio.NewReader(os.Stdin)
//   fmt.Print("Enter text: ")
//   text, _ := reader.ReadString('\n')
//   fmt.Println(text)
//
//   fmt.Println("Enter text: ")
//   text2 := ""
//   fmt.Scanln(text2)
//   fmt.Println(text2)
// }
