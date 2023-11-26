package main

import (
	"fmt"
	"time"
)

type WorkerPool struct {
	concurrency int
	Queue       JobQueue
	Results     OutQueue
}

type JobQueue = chan int
type OutQueue = chan int

func (p *WorkerPool) Run() {
	for i := 0; i < p.concurrency; i++ {
		go func(idx int) {
			for {
				payload := <-p.Queue
				out, err := work_it(idx, payload)
				if err != nil {
					fmt.Println(err)
					continue
				}
				if p.Results != nil {
					p.Results <- out
				}
			}
		}(i)
	}
}

func work_it(wid, payload int) (int, error) {
	fmt.Printf("Worker %d: received payload %d\n", wid, payload)
	time.Sleep(time.Millisecond * 1500)
	fmt.Printf("Worker %d: done\n", wid)
	return payload, nil
}

func produce(queue chan int, done chan bool) {
	defer close(queue)
	for i := 0; i <= 10; i++ {
		queue <- i
	}
	done <- true
}

func main_engine() {
	queue := make(chan int)
	done := make(chan bool)
	results := make(chan int)
	worker_pool := WorkerPool{
		concurrency: 3,
		Queue:       queue,
		Results:     results,
	}
	go func() { worker_pool.Run() }()
	go produce(queue, done)

loop:
	for {
		select {
		case <-worker_pool.Results:
			continue
		case <-done:
			break loop
		}
	}

}
