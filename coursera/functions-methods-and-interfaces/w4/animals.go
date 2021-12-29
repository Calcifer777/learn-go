/*
Write a program which allows the user to create a set of animals and to get
information about those animals. Each animal has a name and can be either a
cow, bird, or snake. With each command, the user can either create a new animal
of one of the three types, or the user can request information about an animal
that he/she has already created. Each animal has a unique name, defined by the
user. Note that the user can define animals of a chosen type, but the types of
animals are restricted to either cow, bird, or snake. The following table
contains the three types of animals and their associated data.

Animal  Food_eaten  Locomtion_method  Spoken_sound
cow     grass       walk              moo
bird    worms       fly               peep
snake   mice        slither           hsss

Your program should present the user with a prompt, >, to indicate that the
user can type a request. Your program should accept one command at a time from
the user, print out a response, and print out a new prompt on a new line. Your
program should continue in this loop forever. Every command from the user must
be either a "newanimal" command or a "query" command.

Each "newanimal" command must be a single line containing three strings. The
first string is "newanimal". The second string is an arbitrary string which will
be the name of the new animal. The third string is the type of the new animal,
either cow, bird, or snake.  Your program should process each newanimal command
by creating the new animal and printing "Created it!" on the screen.

Each "query" command must be a single line containing 3 strings. The first string
is "query". The second string is the name of the animal. The third string is the
name of the information requested about the animal, either eat, move, or speak.
Your program should process each query command by printing out the requested
data.

Define an interface type called Animal which describes the methods of an
animal. Specifically, the Animal interface should contain the methods Eat(),
Move(), and Speak(), which take no arguments and return no values. The Eat()
method should print the animals food, the Move() method should print the
animals locomotion, and the Speak() method should print the animals spoken
sound. Define three types Cow, Bird, and Snake. For each of these three types,
define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and
Snake all satisfy the Animal interface. When the user creates an animal, create
an object of the appropriate type. Your program should call the appropriate
method when the user issues a query command.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (a *Cow) Eat()   { fmt.Println("grass") }
func (a *Cow) Move()  { fmt.Println("walk") }
func (a *Cow) Speak() { fmt.Println("moo") }

type Bird struct{}

func (a *Bird) Eat()   { fmt.Println("worms") }
func (a *Bird) Move()  { fmt.Println("fly") }
func (a *Bird) Speak() { fmt.Println("peep") }

type Snake struct{}

func (a *Snake) Eat()   { fmt.Println("mice") }
func (a *Snake) Move()  { fmt.Println("slither") }
func (a *Snake) Speak() { fmt.Println("hsss") }

func Init(t string) Animal {
	var a Animal
	switch t {
	case "cow":
		a = &Cow{}
	case "bird":
		a = &Bird{}
	case "snake":
		a = &Snake{}
	default:
		fmt.Println("Unknown animal: %s", t)
	}
	fmt.Println("Created it!")
	return a
}

func Query(a Animal, q string) {
	switch q {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	default:
		fmt.Printf("Unknown action: %s\n", q)
	}
}

func main() {
	var animals = make(map[string]Animal)
	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		var line = scanner.Text()
		var splitted = strings.SplitN(line, " ", 3)
		if len(splitted) != 3 {
			fmt.Printf("Can not parse inputs\n")
			continue
		}
		switch splitted[0] {
		case "newanimal":
			animals[splitted[1]] = Init(splitted[2])
		case "query":
			Query(animals[splitted[1]], splitted[2])
		default:
			fmt.Printf("Unknown command: %s\n", splitted[0])
		}
	}
}
