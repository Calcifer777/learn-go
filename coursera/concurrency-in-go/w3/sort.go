/*
Write a program to sort an array of integers. The program should partition the
array into 4 parts, each of which is sorted by a different goroutine. Each
partition should be of approximately equal size. Then the main goroutine should
merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each
goroutine which sorts 1/4 of the array should print the subarray that it will
sort. When sorting is complete, the main goroutine should print the entire
sorted list.

Submission: Upload your source code for the program.
*/

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func ReadInput() []int {
	var arr = make([]int, 0)
	var s string
	var i int
	fmt.Printf("Write a sequence of integers ('X' to stop)\n")
	for {
		fmt.Printf("> ")
		fmt.Scan(&s)
		if s == "X" {
			break
		}
		i, _ = strconv.Atoi(s)
		arr = append(arr, i)
	}
	return arr
}

func split(arr []int) [][]int {
	var chunks = make([][]int, 4)
	var chunk_id int = 0
	for i := 0; i < len(arr); i++ {
		chunks[chunk_id] = append(chunks[chunk_id], arr[i])
		chunk_id = (chunk_id + 1) % 4
	}
	return chunks
}

func MySort(arr []int, c chan []int) {
	fmt.Println("Sorting:", arr)
	sort.Ints(arr)
	c <- arr
}

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

func FourMerge(chunks [][]int) []int {
  var half1 = merge(chunks[0], chunks[1])
  var half2 = merge(chunks[2], chunks[3])
	return merge(half1, half2)
}

func main() {
	var arr = ReadInput()
	var chunks = split(arr)
	var c = make(chan []int, 4)
	for i := 0; i < 4; i++ {
		go MySort(chunks[i], c)
	}
	var sortedChunks = make([][]int, 4)
	for i := 0; i < 4; i++ {
		sortedChunks[i] = <- c
	}
	var sortedArray = FourMerge(sortedChunks)
	fmt.Println("Sorted: ", sortedArray)
}
