package main

//Complete the function solveMeFirst to compute the sum of two integers.

import "fmt"
func solveMeFirst(a uint32, b uint32) uint32 {
    return a + b
}

func simpleArraySum(ar []int32) int32 {
     var x int32 = 0
     for i := 0; i < len(ar); i++ {
            x = ar[i] + x
         }
         return x
}
// Complete the solve function below.
func solve(a []int32, b []int32) []int32 {
    var result []int32 = make([]int32, 2)
    for i := 0; i < len(a); i++ {
        if a[i] > b[i] {
            result[0] = result[0]+ 1
        } else if a[i] == b[i] {
            continue
        } else {
            result[1] = result[1] + 1
        }
    }
    return result
}
func main() {

     morenums := []int32{1, 2, 3, 4, 10, 11}

     fmt.Println(simpleArraySum(morenums))
     numa := []int32{5, 6, 7}
     numb := []int32 {3, 6, 10}
     solve(numa, numb)
}
