package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	NUM_WORKERS  int = 4
	NUM_MESSAGES int = 20
)

func producer(ch chan<- int) {
	defer close(ch)
	for i := 0; i < NUM_MESSAGES; i++ {
		fmt.Printf("Producing %d\n", i)
		ch <- i
		time.Sleep(time.Millisecond * 100)
	}
}

func do_work(worker_id int, payload int) (int, error) {
	fmt.Printf("Consumer %d, received %d\n", worker_id, payload)
	time.Sleep(time.Millisecond * 1200)
	if payload%7 == 0 {
		return 0, fmt.Errorf("error consuming %d", payload)
	} else {
		return payload, nil
	}
}

func consume(queue chan int, outs chan int, errors chan error, wg *sync.WaitGroup) {
	for i := 0; i < NUM_WORKERS; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			for payload := range queue {
				out, err := do_work(idx, payload)
				_, _ = out, err
				if err == nil {
					fmt.Printf("Done with %d\n", payload)
					outs <- out
				} else {
					errors <- err
					fmt.Printf("Error consuming %d\n", payload)
				}
			}
		}(i)
	}
}

func run_spmc() {
	ch_in := make(chan int)
	ch_out := make(chan int)
	ch_err := make(chan error)
	done := make(chan bool)

	wg := sync.WaitGroup{}
	go producer(ch_in)
	consume(ch_in, ch_out, ch_err, &wg)
	go func() {
		wg.Wait()
		done <- true
	}()

	var (
		errs []error
		rsps []int
	)

loop:
	for {
		select {
		case out := <-ch_out:
			rsps = append(rsps, out)
		case err := <-ch_err:
			errs = append(errs, err)
		case <-done:
			break loop
		}
	}

	fmt.Println("Errors", errs)
	fmt.Println("Responses", rsps)

	fmt.Println("Done consuming")
}
