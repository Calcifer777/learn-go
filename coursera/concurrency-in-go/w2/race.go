/*
Write two goroutines which have a race condition when executed concurrently.
Explain what the race condition is and how it can occur.

Upload your source code for the program along with your written explanation of race conditions.
*/

package main

import "time"
import "fmt"

/*
A race condition is a bug in a concurrent program caused by multiple goroutines
trying to access the same resource at the same time. The resource can be some
data, a OS resource (e.g. file), a network resource (e.g. a server), etc.

A data race occurs, for example, when two goroutines access the same variable 
concurrently and at least one of the accesses is a write. This behavior
can lead to crashed, memory corruption, and non-deterministic behavior of the
program.

In this example the two goroutines increase a shared 'counter' variable based
on its value (i.e. try to read and write the 'counter' variable)

Concurrent access to 'counter' means that the outcome of the program is 
non-deterministic.
*/

var counter int = 0

func func1() {
	for i:=0; i<1e6; i++ {
		if i <= counter {
			counter++
		}
	}
	fmt.Printf("Done\n")
}

func main() {
	go func1()
	go func1()
	time.Sleep(1 * time.Second)
	fmt.Printf("Counter value: %d\n", counter)
}
