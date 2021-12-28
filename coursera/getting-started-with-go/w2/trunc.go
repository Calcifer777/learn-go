package main

import "fmt"

func main() {
	var f float64
	fmt.Printf("Write a float: ")
	fmt.Scan(&f)
	fmt.Printf("%0.f\n", f)
}
