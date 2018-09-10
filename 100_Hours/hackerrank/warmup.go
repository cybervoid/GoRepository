package main

import (
	"fmt"
	"strings"
	"strconv"
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

func miniMaxSum(arr []int32) {
	var high int64 = -9223372036854775808
	var low int64 = 9223372036854775807

	for index := 0; index < len(arr); index++ {
		var total int64 = 0
		for i := 0; i < len(arr); i++ {
			if i != index {
				total += int64(arr[i])
			}
		}
		if total > high {
			high = total
		}
		if total < low {
			low = total
		}
	}
	fmt.Print(low)
	fmt.Print(" ")
	fmt.Print(high)
}
func runMiniMaxSum() {
	var arr []int32 = []int32{1, 2, 3, 4, 5}
	miniMaxSum(arr)
}
func timeConversion(s string) string {
	var arr []string = strings.Split(s, ":")
	var hour int64
	var minute int64
	var second int64
    var isPM bool = false
	for i := 0; i < len(arr); i++ {
		if strings.HasSuffix(arr[i], "AM") {
			isPM = false
			arr[i] = strings.TrimSuffix(arr[i], "AM")
		} else if strings.HasSuffix(arr[i], "PM") {
			isPM = true
			arr[i] = strings.TrimSuffix(arr[i], "PM")
		}
		in, _ := strconv.ParseInt(arr[i], 10, 32)

		if i == 0 {
			hour = in
		} else if i == 1 {
			minute = in
		} else if i == 2 {
			second = in
		}
	}
	if isPM == true && hour != 12 {
		hour = hour + 12

	} else if hour == 12 && isPM == false {
		hour = 0
	}
	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
}
func runTimeConversion() {
	var s string = "12:45:54PM"
	fmt.Println("Converting: " + s + "  ...")
	time := timeConversion(s)
	fmt.Println(time)
}
func gradingStudents(grades []int32) []int32 {
	var newGrades []int32 = make([]int32, 0)
	for i := 0; i < len(grades); i++ {
		var remainder = grades[i] % 5
		var rounded bool = false
		if remainder >= 3 {
			if grades[i] >= 38 {
				rounded = true
				var input int32 = grades[i] + (5 - remainder)
				newGrades = append(newGrades, input)
			}
		}
		if rounded == false {
			newGrades = append(newGrades, grades[i])
		}
	}
	return newGrades
}
func runGradingStudents() {
	inputs := []int32{ 73, 67, 38, 33 }
	results := gradingStudents(inputs)
	for i := 0; i < len(results); i++ {
		fmt.Println(results[i])
	}
}
func main() {
	runGradingStudents()
	//runTimeConversion()
	//runMiniMaxSum()
	//runStaircase()
	//runPlusMinus()
	//runDiagnalDifference()
}
