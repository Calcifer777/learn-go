package day3

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"unicode"
)

type Num struct {
	N      int
	Xstart int
	Xend   int
	Y      int
}

func NewNumber(chars []rune, xStart int, Xend int, y int) (*Num, error) {
	n, e := strconv.Atoi(string(chars))
	if e != nil {
		return nil, e
	}
	return &Num{n, xStart, Xend, y}, nil
}

type Sym struct {
	V string
	X int
	Y int
}

func ParseFile(f *os.File) ([]Num, []Sym, error) {
	buf := bufio.NewScanner(f)
	acc := make([]rune, 0)
	var xStart, numLen int
	numbers := make([]Num, 0)
	symbols := make([]Sym, 0)
	row := 0
	for buf.Scan() {
		line := buf.Text()
		slog.Debug(line)
		xStart, numLen = -1, 0
		for col, ch := range line {
			if unicode.IsDigit(ch) {
				if len(acc) == 0 {
					xStart = col
				}
				numLen += 1
				acc = append(acc, ch)
			} else {
				if len(acc) > 0 {
					newN, e := NewNumber(acc, xStart, xStart+numLen, row)
					if e != nil {
						return nil, nil, e
					}
					numbers = append(numbers, *newN)
					xStart, numLen = -1, 0
					acc = make([]rune, 0)
				}
				if ch == '.' {
					continue
				} else {
					symbols = append(symbols, Sym{string(ch), col, row})
				}
			}
		}
		if len(acc) > 0 {
			newN, e := NewNumber(acc, xStart, xStart+numLen, row)
			if e != nil {
				return nil, nil, e
			}
			numbers = append(numbers, *newN)
			xStart, numLen = -1, 0
			acc = make([]rune, 0)
		}
		row += 1
	}
	return numbers, symbols, nil
}

func filterN(n Num, syms []Sym) bool {
	for _, s := range syms {
		checkX := (n.Xend+1 >= s.X) && (n.Xstart-1 <= s.X)
		checkY := (n.Y >= s.Y-1) && (n.Y <= s.Y+1)
		slog.Debug(fmt.Sprintf("%v - %v - X: %v, Y: %v", n, s, checkX, checkY))
		if checkX && checkY {
			return true
		}
	}
	return false
}

func filterNs(ns []Num, syms []Sym) []Num {
	res := make([]Num, 0)
	for _, n := range ns {
		if filterN(n, syms) {
			res = append(res, n)
		}
	}
	return res
}

func Part1(path string) (int, error) {
	f, e := os.Open(path)
	if e != nil {
		slog.Error(fmt.Sprintf("Cound not open file at %s", path))
		return -1, e
	}
	defer f.Close()
	ns, syms, e := ParseFile(f)
	if e != nil {
		slog.Error(fmt.Sprintf("Cound not parse file: %v", e))
		return -1, e
	}
	for _, n := range ns {
		slog.Debug(fmt.Sprintf("%v", n))
	}
	for _, s := range syms {
		slog.Debug(fmt.Sprintf("%v", s))
	}
	sel := filterNs(ns, syms)
	sum := 0
	for _, n := range sel {
		slog.Debug(fmt.Sprintf("%v", n))
		sum += n.N
	}
	return sum, nil
}
