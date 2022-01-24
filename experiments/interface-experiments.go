package main

import (
	"fmt"
)

type Item struct {
	name string
}

type Fooable interface {
	foo()
}

func (a Item) foo() {
	fmt.Printf("Hello from %s\n", a.name)
}

func DoFoo1(i Item) {
	i.foo()
}

func DoFoo2[T Fooable](i *T) {
	(*i).foo()
}

// Exp 2

type Speaker interface {
	speak()
}

type Dog struct {
	name string
}

func (d *Dog) speak() {
	fmt.Printf("Hello from dog %s\n", d.name)
}

type Cat struct {
	name string
}

func (c *Cat) speak() {
	fmt.Printf("Hello from cat %s\n", c.name)
}

func Speak(s Speaker) {
	switch v := s.(type) {
	case *Dog:
		v.speak()
	case *Cat:
		v.speak()
	}
}

func Speak2[T Speaker](t T) {
	var i Speaker = t
	switch v := i.(type) {
	case *Dog:
		v.speak()
	case *Cat:
		v.speak()
	}
}

func main() {
	fmt.Printf("Hello\n")
	// Exp 1: Interface with generics
	i := Item{"gladys"}
	i.foo()
	DoFoo2(&i)
	// Exp 2: switch case on type
	d := Dog{"barky"}
	c := Cat{"spotty"}
	Speak(&d)
	Speak2(&c)
}
