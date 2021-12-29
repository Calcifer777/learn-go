/*
Write a program which allows the user to get information about a predefined set
of animals. Three animals are predefined, cow, bird, and snake. Each animal can
eat, move, and speak.

The user can issue a request to find out one of three things about an animal:
1) the food that it eats
2) its method of locomotion,
3) the sound it makes when it speaks.

The following table contains the three animals and their associated data which
should be hard-coded into your program.

Animal  Food_eaten  Locomotion_method  Spoken_sound
cow     grass       walk               moo
bird    worms       fly                peep
snake   mice        slither            hsss

Your program should present the user with a prompt, >, to indicate that the
user can type a request. Your program accepts one request at a time from the
user, prints out the answer to the request, and prints out a new prompt. Your
program should continue in this loop forever. Every request from the user must
be a single line containing 2 strings. The first string is the name of an
animal, either cow, bird, or snake. The second string is the name of the
information requested about the animal, either eat, move, or speak. Your
program should process each request by printing out the requested data.

You will need a data structure to hold the information about each animal. Make
a type called Animal which is a struct containing three fields:food,
locomotion, and noise, all of which are strings. Make three methods called
Eat(), Move(), and Speak(). The receiver type of all of your methods should be
your Animal type. The Eat() method should print the animals food, the Move()
method should print the animals locomotion, and the Speak() method should print
the animals spoken sound. Your program should call the appropriate method when
the user makes a request.

Submit your Go program source code.

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat()   { fmt.Println(a.food) }
func (a Animal) Move()  { fmt.Println(a.locomotion) }
func (a Animal) Speak() { fmt.Println(a.noise) }

func main() {
	var cow = Animal{"grass", "walk", "moo"}
	var bird = Animal{"worms", "fly", "peep"}
	var snake = Animal{"mice", "slither", "hsss"}
	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		var line = scanner.Text()
		var splitted = strings.SplitN(line, " ", 2)
		if len(splitted) != 2 {
			fmt.Printf("Can not parse inputs\n")
			continue
		}
		var animal, action = splitted[0], splitted[1]
		switch animal {
		case "cow":
			switch action {
			case "eat":
				cow.Eat()
			case "move":
				cow.Move()
			case "speak":
				cow.Speak()
			default:
				fmt.Printf("Unknown action: %s\n", action)
			}
		case "bird":
			switch action {
			case "eat":
				bird.Eat()
			case "move":
				bird.Move()
			case "speak":
				bird.Speak()
			default:
				fmt.Printf("Unknown action: %s\n", action)
			}
		case "snake":
			switch action {
			case "eat":
				snake.Eat()
			case "move":
				snake.Move()
			case "speak":
				snake.Speak()
			default:
				fmt.Printf("Unknown action: %s\n", action)
			}
		default:
			fmt.Printf("Unknown animal: %s\n", animal)
		}
	}
}
