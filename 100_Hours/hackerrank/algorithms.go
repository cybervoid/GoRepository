package main

import (
	"fmt"
)

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
    var difference int32 = 0
    for i := int32(1); i < 10000; i++ {
        var k1 int32 = x1 + (v1 * i)
        var k2 int32 = x2 + (v2 * i)
        if k1 == k2 {
            return "YES"
        }
        var newDifference int32 = 0
        if k1 > k2 {
            newDifference = k1 - k2
        } else {
            newDifference = k2 - k1
        }

        if difference != 0 {
            if newDifference > difference {
                return "NO"
            }
        }
        difference = newDifference
    }
    return "NO"
}
func getTotalX(a []int32, b []int32) int32 {
    var total int32 = 0
    

}
func runGetTotalX() {
    var a []int32 = []int32 { 2, 4 }
    var b []int32 = []int32 { 16, 32, 96 }
    total := getTotalX(a, b)
}

func runKangaroo() {
    fmt.Println(kangaroo(0, 2, 5, 3))
}
func main() {
    runKangaroo()
}
