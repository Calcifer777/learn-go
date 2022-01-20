package main

import (
  "testing"
)

func TestMean(t *testing.T) {
  arr := []int{1,2,3,4,5}
  want := 3
  mean := Mean(arr)
  if mean != want {
    t.Fatalf("Mean([]int{1,2,3,4,5}) = %d, want 3", mean)
  }
}

func TestMedian(t *testing.T) {
  arr := []int{1,2,3,4,5}
  want := 3
  median := Median(arr)
  if median != want {
    t.Fatalf("Median([]int{1,2,3,4,5}) = %d, want 3", median)
  }
}

func TestMedianEvenSize(t *testing.T) {
  arr := []int{1,2,4,5}
  want := 3
  median := Median(arr)
  if median != want {
    t.Fatalf("Median([]int{1,2,4,5}) = %d, want 3", median)
  }
}
