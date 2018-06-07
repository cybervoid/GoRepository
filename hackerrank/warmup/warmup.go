package main

import (
	"fmt"
)

func solveMeFirst(a uint32, b uint32) uint32 {
	return a + b
}

func aVeryBigSum(ar []int64) int64 {
	var big int64

	for i := 0; i < len(ar); i++ {
		big = ar[i] + big
	}
	return big
}
func simpleArraySum(ar []int32) int32 {
	var x int32 = 0
	for i := 0; i < len(ar); i++ {
		x = ar[i] + x
	}
	return x
}
func diagonalDifference(a [][]int32) int32 {
	var primaryDiag int32 = 0

	var rows = 0
	for i, v1 := range a {
		primaryDiag = primaryDiag + v1[i]
		rows = i
	}

	var secondaryDiag int32 = 0
	var k = rows
	for s := 0; s <= rows; s++ {
		fmt.Printf("\t [%d] [%d] ", k, s)
		fmt.Println(a[k][s])
		secondaryDiag = a[k][s] + secondaryDiag

		k--
	}
	total := primaryDiag - secondaryDiag
	if total < 0 {
		total = total * -1
	}

	return total
}
func runDiagnalDifference() {
	multi := [][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}
	//fmt.Printf("%t", multi)
	fmt.Println(diagonalDifference(multi))
}
func main() {

	runDiagnalDifference()
}
