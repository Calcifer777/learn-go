/*
Write a Bubble Sort program in Go. The program should prompt the user to type
in a sequence of up to 10 integers. The program should print the integers out
on one line, in sorted order, from least to greatest. Use your favorite search
tool to find a description of how the bubble sort algorithm works.

As part of this program, you should write a function called BubbleSort() which
takes a slice of integers as an argument and returns nothing. The BubbleSort()
function should modify the slice so that the elements are in sorted order.

A recurring operation in the bubble sort algorithm is the Swap operation which
swaps the position of two adjacent elements in the slice. You should write a
Swap() function which performs this operation. Your Swap() function should take
two arguments, a slice of integers and an index value i which indicates a
position in the slice. The Swap() function should return nothing, but it should
swap the contents of the slice in position i with the contents in position i+1.

Submit your Go program source code.
*/

package main

import (
	"fmt"
  "sort"
	"strconv"
)

func ReadInput() []int {
	fmt.Println("Enter a list of integers (enter 'X' to stop):")
	var s string
	var n int
	var ns = make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		fmt.Scan(&s)
		if s == "X" {
			break
		}
		n, _ = strconv.Atoi(s)
		ns = append(ns, n)
	}
	return ns
}

func Swap(ns []int, idx int) {
	var tmp = ns[idx]
	ns[idx] = ns[idx+1]
	ns[idx+1] = tmp
}

func BubbleSort(ns []int) {
	var flag bool
	var l = len(ns)
	for {
		flag = true
		for i := 0; i < l-1; i++ {
			if ns[i] > ns[i+1] {
				Swap(ns, i)
				flag = false
			}
		}
		if flag {
			break
		}
		l--
	}
}

func main() {
	ns := ReadInput()
	BubbleSort(ns)
	fmt.Print("The sorted array is: ", ns)
}
