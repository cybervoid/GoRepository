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

// Complete the plusMinus function below.
func plusMinus(arr []int32) {
	var zero = 0
	var positive = 0
	var negative = 0
	var total = 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			negative++
		}
		if arr[i] == 0 {
			zero++
		}
		if arr[i] > 0 {
			positive++
		}
		total = i + 1
	}
	var percentPositive float32 = 0
	if positive != 0 {
		percentPositive = float32(positive) / float32(total)
	}
	var percentZero float32 = 0
	if zero != 0 {
		percentZero = float32(zero) / float32(total)
	}
	var percentNegative float32 = 0
	if negative != 0 {
		percentNegative = float32(negative) / float32(total)
	}
	fmt.Println(percentPositive)
	fmt.Println(percentZero)
	fmt.Println(percentNegative)
}
func runPlusMinus() {
	var arr []int32 = []int32{-4, 3, -9, 0, 4, 1}
	plusMinus(arr)
}

func staircase(n int32) {
	var i int32 = 0
	for i = 0; i < n; i++ {
		var output string = ""
		var j int32 = 0
		for j = 0; j < n; j++ {
			if i < j {
				output = " " + output
			} else {
				output = "#" + output
			}
		}
		fmt.Println(output)
	}
}

func runStaircase() {
	staircase(4)
}
func main() {
	runStaircase()
	//runPlusMinus()
	//runDiagnalDifference()
}
