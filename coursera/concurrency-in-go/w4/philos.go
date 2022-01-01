/*
Implement the dining philosopher’s problem with the following
constraints/modifications.
[X] - There should be 5 philosophers sharing chopsticks, with one chopstick between
each adjacent pair of philosophers.
[X] - Each philosopher should eat only 3 times (not in an infinite loop as we did in
lecture)
[X] - The philosophers pick up the chopsticks in any order, not lowest-numbered first
(which we did in lecture).
[ ] - In order to eat, a philosopher must get permission from a host which executes
in its own goroutine.
[ ] - The host allows no more than 2 philosophers to eat concurrently.
[X] - Each philosopher is numbered, 1 through 5.
[X] - When a philosopher starts eating (after it has obtained necessary locks) it
prints “starting to eat <number>” on a line by itself, where <number> is the
number of the philosopher.
[X] - When a philosopher finishes eating (before it has released its locks) it prints
“finishing eating <number>” on a line by itself, where <number> is the number
of the philosopher.
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	philos_num      int = 5
	chops_num       int = 5
	eat_num         int = 3
	eat_concurrency int = 2
)

type Philo struct {
	id           int
	lChop, rChop *Chop
	eatCounter   int
	waiting      bool
}

type Chop struct {
	mut sync.Mutex
}

type Host struct {
	eatingPhilos []bool
	num_eating   int
	num_servings int
}

func mod(d, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}

func (host *Host) Schedule(chAskEat chan int, chAllowEat chan int, chDoneEat chan int, wg *sync.WaitGroup) {
	// Counter for the number of times a Philo eats
	for {
		// All Philos are done eating
		select {
		case philoId := <-chDoneEat:
			if !host.eatingPhilos[philoId] {
				fmt.Printf("[HOST] %d finished eating when never allowed", philoId)
				panic("Unreachable")
			}
			fmt.Printf("[HOST] %d finished eating\n", philoId)
			host.eatingPhilos[philoId] = false
			host.num_eating--
			host.num_servings++
			fmt.Printf("[HOST] Num eating philos: %d\n", host.num_eating)
		case philoId := <-chAskEat:
			var condAllow = host.num_eating < 2 &&
				!host.eatingPhilos[philoId] &&
				!host.eatingPhilos[mod(philoId+1, 5)] &&
				!host.eatingPhilos[mod(philoId-1, 5)]
			if condAllow {
				host.eatingPhilos[philoId] = true
				host.num_eating++
				chAllowEat <- philoId
				fmt.Printf("[HOST] %d asked eating: allowed\n", philoId)
				fmt.Printf("[HOST] Num eating philos: %d\n", host.num_eating)
			} else {
				fmt.Printf("[HOST] %d asked eating: bounced\n", philoId)
				chAskEat <- philoId
			}
		default:
			fmt.Printf("[HOST] Num servings: %d\n", host.num_servings)
		}
		if host.num_servings == 15 {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	wg.Done()
}

func (philo *Philo) Eat() {
	// philo.lChop.mut.Lock()
	// philo.rChop.mut.Lock()
	fmt.Printf("[%d] starting to eat\n", philo.id)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Printf("[%d] finishing eating\n", philo.id)
	// philo.lChop.mut.Unlock()
	// philo.rChop.mut.Unlock()
	philo.eatCounter++
}

func (philo *Philo) Dine(chAskEat chan int, chAllowEat chan int, chDoneEat chan int, wg *sync.WaitGroup) {
	for {
		if !philo.waiting {
			// Ask for turn
			chAskEat <- philo.id
			philo.waiting = true
		} else {
			// Check permission to eat
			allowedId := <-chAllowEat
			if allowedId == philo.id {
				// If my turn, eat
				philo.Eat()
				chDoneEat <- philo.id
				philo.waiting = false
			} else {
				// ... return the allowed id
				chAllowEat <- allowedId
			}
		}
		if philo.eatCounter >= eat_num {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	wg.Done()
}

func main() {
	// Init
	var philos = make([]*Philo, 5)
	var chops = make([]*Chop, 5)
	for i := 0; i < chops_num; i++ {
		chops[i] = new(Chop)
	}
	for i := 0; i < philos_num; i++ {
		philos[i] = &Philo{i, chops[i], chops[(i+1)%5], 0, false}
	}
	var wg sync.WaitGroup
	var chAskEat = make(chan int, 100)
	var chAllowEat = make(chan int, 100)
	var chDoneEat = make(chan int, 100)
	// Start Host
	wg.Add(1)
	var host = Host{make([]bool, 5), 0, 0}
	go host.Schedule(chAskEat, chAllowEat, chDoneEat, &wg)
	// Dine
	for i := 0; i < philos_num; i++ {
		wg.Add(1)
		go philos[i].Dine(chAskEat, chAllowEat, chDoneEat, &wg)
	}
	wg.Wait()
	fmt.Printf("[Done]. Num eating philos: %d\n", host.num_eating)
}
