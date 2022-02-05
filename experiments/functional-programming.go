package main

import (
	"fmt"
)

func Map[T any, V any](arr []T, f func(t T) V) []V {
	result := make([]V, len(arr))
	for idx, x := range arr {
		result[idx] = f(x)
	}
	return result
}

func Compose[A any, B any, C any](g func(b B) C, f func(a A) B) func(a A) C {
	return func(a A) C {
		return g(f(a))
	}
}

func Filter[T any](arr []T, f func(t T) bool) []T {
	result := make([]T, 0)
	for _, t := range arr {
		if f(t) {
			result = append(result, t)
		}
	}
	return result
}

// Doesn't work when imported from a different package
func Zip[A any, B any](ts []A, vs []B) []struct{ _1 A; _2 B } {
  pairs := make([]struct{_1 A; _2 B }, 0)
  if len(ts) != len(vs) {
    panic("Different slice lengths")
  }
  for i:=0; i<len(ts); i++ {
    pair := struct {_1 A; _2 B}{ ts[i], vs[i] }
    pairs = append(pairs, pair)
  }
  return pairs
}

func main() {
	fmt.Printf("Hello world\n")
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	add1 := func(x int) int { return x + 1 }
	double := func(x int) int { return x * 2 }
	isEven := func(x int) bool { return x % 2 == 0 }
	arr2 := Map(arr, add1)
	for _, x := range arr2 {
		fmt.Printf("%d\n", x)
	}
	fmt.Printf("%d\n", Compose(double, add1)(7))
	fmt.Printf("%v\n", Filter(arr, isEven))
}
