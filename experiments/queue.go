package main

import (
  "fmt"
)

// LIFO queue implementation
type Queue struct {
  head *int
  tail *Queue
}

func (q Queue) Pop() (Queue, *int) {
  if q.tail != nil {
    return Queue{q.tail.head, q.tail.tail}, q.head
  } else {
    return Queue{nil, nil}, q.head
  }
}

func (q Queue) Push(v int) Queue {
  return Queue{head: &v, tail: &q}
}

func (q Queue) HasNext() bool {
  return q.head != nil
}

func main() {
  x := 0
  queue := Queue{&x, nil}
  for i := 1; i < 3; i++ {
    queue = queue.Push(i)
  }
  var v *int
  for {
    if !queue.HasNext() {
      break
    }
    queue, v = queue.Pop()
    fmt.Printf("Popping %d\n", *v)
  }
  for i := 0; i < 10; i++ {
    queue = queue.Push(i)
  }
  for {
    if !queue.HasNext() {
      break
    }
    queue, v = queue.Pop()
    fmt.Printf("Popping %d\n", *v)
  }
}
