package main


import (
  "fmt"
  "math"
  "strings"
  "strconv"
  "utils"
)

func Sort(arr []int) []int {
  size := len(arr)
  for {
    flag := true
    for idx, _ := range arr[:size-1] {
      if arr[idx] > arr[idx+1] {
        tmp := arr[idx]
        arr[idx] = arr[idx+1]
        arr[idx+1] = tmp
        flag = false
      }
    }
    if flag {
      break
    }
  }
  return arr
}

func ParseInput(s string) []int {
  xs := strings.Split(s, ",")
  arr := make([]int, 0)
  for _, x := range xs {
    i, _ := strconv.Atoi(x)
    arr = append(arr, i)
  }
  return arr
}

func Median(arr []int) int {
  sorted := Sort(arr)
  size := len(arr)
  if size % 2 == 0 {
    return (sorted[size/2-1] + sorted[size/2]) / 2
  } else {
    return sorted[size/2+1]
  }
}

func Mean(arr []int) int {
  size := len(arr)
  sum := 0
  for _, i := range arr {
    sum += i
  }
  res := float64(sum)/float64(size)
  fmt.Println(res)
  return int(math.Floor(res))
}

func abs(i int) int {
  if i > 0 {
    return i
  } else {
    return -i
  }
}

func SumResid(arr []int, median int) int {
  sum := 0
  for _, i := range(arr) {
    sum += abs(i - median)
  }
  return sum
}

func SumResid2(arr []int, median int) int {
  sum := 0
  for _, i := range(arr) {
    sum += Fuel(i, median)
  }
  return sum
}

func Fuel(start int, end int) int {
  diff := abs(start - end)
  fuel := 0
  for i := 0; i <= diff; i++ {
    fuel += i
  }
  return fuel
}

func Part1() {
  input, _ := utils.ReadLines("input.txt")
  arr := ParseInput(input[0])
  sorted := Sort(arr)
  median := Median(arr)
  fuel := SumResid(arr, median)
  fmt.Printf("%v\n", sorted)
  fmt.Printf("Median: %v\n", median)
  fmt.Printf("Sum of resids: %d\n", fuel)
}

func Part2() {
  input, _ := utils.ReadLines("input.txt")
  arr := ParseInput(input[0])
  sorted := Sort(arr)
  mean := Mean(arr)
  fuel := SumResid2(arr, mean)
  fmt.Printf("%v\n", sorted)
  fmt.Printf("Mean: %v\n", mean)
  fmt.Printf("Sum of resids: %d\n", fuel)
}

func main() {
  Part1()
  Part2()
}
