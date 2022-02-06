package main

import (
  "fmt"
  "math/rand"
  "time"
  "sync"
)

func Produce(c chan int, wg *sync.WaitGroup) {
  defer wg.Done()
  for i := 0; i < 10; i++ {
    fmt.Printf("Sending: %d\n", i)
    c <- i
    time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
  }
  close(c)
}

func Consume(c chan int, wg *sync.WaitGroup) {
  defer wg.Done()
  for i := range c {
    fmt.Printf("Received: %d\n", i)
    time.Sleep(time.Millisecond * time.Duration(rand.Intn(1500)))
  }
}

func main() {
  var wg sync.WaitGroup
  wg.Add(2)
  ns := make(chan int, 10)
  go Produce(ns, &wg)
  go Consume(ns, &wg)
  wg.Wait()
}
