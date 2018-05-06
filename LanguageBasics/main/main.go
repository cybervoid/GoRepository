package main

import (
  "../basics"
  "../examples"
  "fmt"
  "log"
  "sort"
)

func main() {
  m := getOptions()
  var keys []int
  for k := range m {
    keys = append(keys, k)
  }
  sort.Ints(keys)

  for _, k := range keys {
    fmt.Printf("code: %v \tRuns:\t%v\n",k, m[k] )
  }
  fmt.Printf("Enter Start Code:")
  var input int
  _, err := fmt.Scanf("%d", &input)

  if(HandleError(err, 1) == false) {
    runExample(input, m[input])
  }
}

func HandleError(err error, catchHandler int) bool {
  if (err) != nil {
    fmt.Println(err)
    if(catchHandler == 0) {
      return true
    }
    if(catchHandler == 1) {
      log.Fatal(err)
      return true
    }
  } else {
    return false
  }
  return false
}

func runExample(selected int, exName string) {
  switch selected {
    case 0:
      basics.ShortHand()
    case 1:
      basics.AsVars()
    case 2:
      basics.Closure()
    case 3:
      basics.BlankIdentifierExample()
    case 4:
      basics.MemAddress()
    case 5:
      basics.UseMemAddress()
    case 6:
      basics.PointerExample()
    case 7:
      basics.UseMemAddressAgain()
    case 8:
      fmt.Println("For Loop")
      basics.ForLoop()
      fmt.Println("While Loop")
      basics.WhileLoop()
    case 9:
      basics.Cast()
    case 10:
      basics.RuneCounter()
    case 11:
      basics.SwitchFallthroughStatements()
    case 12:
      fmt.Println("Switch:\tint")
      basics.SwitchOnType(5)
      fmt.Println("Switch:\tstring")
      basics.SwitchOnType("testing")
      fmt.Println("Switch:\tstruct")
      var t = basics.Contact{ "Hello!", "Ryan!" }
      basics.SwitchOnType(t)
    case 13:
      basics.IfInitializationStatements(false)
    case 14:
      basics.GreetFull("Bill", "Burr")
    case 15:
      fmt.Println("Basic Return Example: ")
      fmt.Println(basics.ReturnExample("this is Ryan"))
      fmt.Println("Assigned Returns: ")
      fmt.Println(basics.ReturnAssigned("Ryan ", "Rules"))
      fmt.Println(basics.ReturnAssigned("Ryan ", "Rules"))
      fmt.Println("Multiple Returns: ")
      fmt.Println(basics.ReturnMultiple("Ryan", "Dubs"))
    case 16:
      fmt.Println(basics.FuncAverageVariadicParams(5.554, 76.766, 123.674, 22.32223))
      basics.FuncAverageVariadicArgs()
    case 17:
      basics.ExpressionExample()
    case 18:
      basics.ClosureExpression()
      basics.ClosureFunc()
      basics.ClosureNotAnon()
    case 19:
      basics.RunCallbacks()
    case 20:
      examples.RunFilter()
    case 21:
      basics.RunRecursiveFactorial()
    case 22:
      basics.RunDefer(false)
    case 23:
      fmt.Println("Pass by Value")
      basics.RunPassByValue()
      fmt.Println("Pass by Reference")
      basics.RunPassByReference()
      fmt.Println("Pass by Reference Map")
      basics.RunPassByReference()
    case 24:
      basics.RunAnonymousSelfExecuting()
    case 25:
      fmt.Println("Slice example")
      basics.SliceExample()
      fmt.Println("Map example")
      basics.MapExample()
    case 26:
      basics.RunInterfacesExample()
    case 27:
      basics.RunParallel()
    case 28:
      fmt.Println("Basic Concurrency")
      basics.RunConcurrency()
      fmt.Println("Race Conditions")
      fmt.Println("run main.go with -race parameter")
      basics.RaceCondition()
    case 29:
      fmt.Println("MutexExample()")
      basics.MutexExample()
    case 30:
      fmt.Println("AtomicityExample()")
      basics.AtomicityExample()
    case 31:
      fmt.Println("Channels - Unbuffered Basic")
      basics.ChannelsBasicUnbufferedExample()
      fmt.Println("Channels - Unbuffered with Close")
      basics.ChannelsBasicUnbufferedExample()
      fmt.Println("Channels - Multiple Unbuffered")
      basics.MultipleChannelsUnbufferedExample()
    case 32:
      fmt.Println("Semaphore - Example 1")
      basics.SemaphoreExample()
    case 33:
      fmt.Println("Structs - Example 1")
      basics.StructsExample()
    default:
      fmt.Println("Invalid option")
  }
  fmt.Println("Completed running: \"" + exName +"\"")
}

func getOptions() (opt map[int] string) {
  opt = make(map[int] string)
  opt[0] = "Short Hand"
  opt[1] ="As Vars"
  opt[2] = "Closure"
  opt[3] = "Blank Identifier"
  opt[4] = "Memory Address"
  opt[5] = "Use Memory Addresses"
  opt[6] = "Pointers"
  opt[7] = "Use Mem Address Again"
  opt[8] ="Loops"
  opt[9] = "Casting"
  opt[10] = "Runes"
  opt[11] = "Switch Fall through Statements"
  opt[12] ="SwitchOnType(interface{})"
  opt[13] ="If Initialization Statements"
  opt[14] ="GreetFull Bill"
  opt[15] = "Return Examples"
  opt[16] = "Variadics"
  opt[17] = "Expression Examples"
  opt[18] = "Closures"
  opt[19] = "Callbacks"
  opt[20] = "Filter(s)"
  opt[21] = "Recursive Factorial"
  opt[22] = "Run defer"
  opt[23] = "Passing values by Value"
  opt[24] = "Anon self executing funcs"
  opt[25] = "Data Structures"
  opt[26] = "Interfaces Example"
  opt[27] = "Parallel Examples"
  opt[28] = "Concurrency Examples"
  opt[29] = "Mutex"
  opt[30] = "Atomicity"
  opt[31] = "Channels"
  opt[32] = "Semaphore Pattern"
  opt[33] = "Structs Examples"
  return opt
}
