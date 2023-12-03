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
	v        int
	colStart int
	colEnd   int
	row      int
}

func NewNumber(chars []rune, colStart int, colEnd int, row int) (*Num, error) {
	n, e := strconv.Atoi(string(chars))
	if e != nil {
		return nil, e
	}
	return &Num{n, colStart, colEnd, row}, nil
}

type Sym struct {
	v   string
	col int
	row int
}

func ParseFile(f *os.File) ([]Num, []Sym, error) {
	buf := bufio.NewScanner(f)
	acc := make([]rune, 0)
	numbers := make([]Num, 0)
	symbols := make([]Sym, 0)
	row := 0
	for buf.Scan() {
		line := buf.Text()
		slog.Debug(line)
		for col, ch := range line {
			if unicode.IsDigit(ch) {
				acc = append(acc, ch)
			} else {
				if len(acc) > 0 {
					newN, e := NewNumber(acc, col-len(acc), col, row)
					if e != nil {
						return nil, nil, e
					}
					numbers = append(numbers, *newN)
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
			col := len(line) - 1
			newN, e := NewNumber(acc, col-len(acc), col, row)
			if e != nil {
				return nil, nil, e
			}
			numbers = append(numbers, *newN)
			acc = make([]rune, 0)
		}
		row += 1
	}
	return numbers, symbols, nil
}

func filterN(n Num, syms []Sym) bool {
	for _, s := range syms {
		checkX := (n.colEnd+1 >= s.col) && (n.colStart-1 <= s.col)
		checkY := (n.row >= s.row-1) && (n.row <= s.row+1)
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
		sum += n.v
	}
	return sum, nil
}
