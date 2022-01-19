package main


import (
  "fmt"
  "strings"
  "strconv"
  "utils"
)

type LanternFish int

type School map[int]int

func ParseInput(s string) School {
  values := strings.Split(s, ",")
  school := make(School, 0)
  for _, v := range values {
    i, _ := strconv.Atoi(v)
    school[i] += 1
  }
  return school
}

func NextDay(school School) School {
  nextDaySchool := make(School)
  for i := 8; i >= 0; i-- {
    v, ok := school[i]
    if !ok {
      continue
    }
    if i > 0 {
      nextDaySchool[i-1] += v
    } else {
      nextDaySchool[6] += v
      nextDaySchool[8] += v
    }
  }
  return nextDaySchool
}

func ToDay(s School, n int) School {
  res := make(School)
  // copy
  for k, v := range s {
      res[k] = v
    }
  for i := 0; i < n; i++ {
    res = NextDay(res)
  }
  return res
}

func SchoolSize(s School) int {
  size := 0
  for _, v := range s {
    size += v
  }
  return size
}

func Part1() {
  input, _ := utils.ReadLines("input.txt")
  school := ParseInput(input[0])
  school = ToDay(school, 80)
  size := SchoolSize(school)
  fmt.Printf("Part 1 -> School Size: %d\n", size)
}

func Part2() {
  input, _ := utils.ReadLines("input.txt")
  school := ParseInput(input[0])
  school = ToDay(school, 256)
  size := SchoolSize(school)
  fmt.Printf("Part 2 -> School Size: %d\n", size)
}

func main() {
  Part1()
  Part2()
}
