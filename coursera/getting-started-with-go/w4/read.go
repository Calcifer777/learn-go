/*
Write a program which reads information from a file and represents it in a
slice of structs. Assume that there is a text file which contains a series of
names. Each line of the text file has a first name and a last name, in that
order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the
first name, and lname for the last name. Each field will be a string of size 20
(characters).

Your program should prompt the user for the name of the text file. Your program
will successively read each line of the text file and create a struct which
contains the first and last names found in the file. Each struct created will
be added to a slice, and after all lines have been read from the file, your
program will have a slice containing one struct for each line in the file.
After reading all lines from the file, your program should iterate through your
slice of structs and print the first and last names found in each struct.

Submit your source code for the program, read.go.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const maxNameLength = 20

type Name struct {
	fname string
	lname string
}

func (name *Name) Set(fname string, lname string) {
	var rs []rune

	name.fname = fname
	if len(fname) > maxNameLength {
		rs = []rune(fname)
		name.fname = string(rs[:maxNameLength])
	} else {
	  name.fname = fmt.Sprintf("%-20s", fname)
	}

	name.lname = lname
	if len(lname) > maxNameLength {
		rs = []rune(lname)
		name.lname = string(rs[:maxNameLength])
	} else {
	  name.lname = fmt.Sprintf("%-20s", lname)
	}
}

func main() {
	var fileName string
	fmt.Print("Enter the file name: ")
	fmt.Scan(&fileName)

	// Read file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var n = new(Name)
	var ns = make([]Name, 0)
	for scanner.Scan() {
		line := scanner.Text()
		splitted := strings.SplitN(line, " ", 2)
		n.Set(splitted[0], splitted[1])
		ns = append(ns, *n)
	}

	// Print names
	for idx, n := range ns {
		fmt.Printf("%d: %s, %s\n", idx, n.fname, n.lname)
	}
}
