package examples

import "fmt"

func filter(numbers []int, callback func(int) bool) []int{
  xs := []int{}
  for _, n := range numbers {
    if callback(n) {
      xs = append(xs, n)
    }
  }
  return xs
}

func RunFilter() {
  xs := filter([]int {1, 2, 3, 4, 15, 20, 30 }, func(n int) bool {
    return n > 10
  })
  fmt.Println(xs) //[2 3 4]
}
