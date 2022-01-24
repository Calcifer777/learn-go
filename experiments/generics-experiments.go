package main

import (
	"fmt"
	"sort"
)

// Option 1
func GenAdd1[T int | float64](arr []T) T {
	var sum T
	for _, v := range arr {
		sum += v
	}
	return sum
}

// Option 2
type Numeric interface {
	int64 | float64
}

func GenAdd2[T Numeric](arr []T) T {
	var sum T
	for _, v := range arr {
		sum += v
	}
	return sum
}

func GenAdd3[T sort.Interface](arr []T) {
	fmt.Printf("Not implemented, yet")
}

func main() {
	// anonymous function
	greet := func() {
		fmt.Printf("Hello world\n")
	}
	greet()
	// Generics
	arr1 := []int{1, 2, 3}
	arr2 := []float64{1.0, 2.0, 3.0}
	fmt.Printf("%v\n", GenAdd1(arr1))
	fmt.Printf("%v\n", GenAdd2(arr2))
	// interface
}
