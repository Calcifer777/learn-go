package day13

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
)

func Part1(path string) (int, error) {
	f, e := os.Open(path)
	if e != nil {
		slog.Error(fmt.Sprintf("Cound not open file at %s", path))
		return -1, e
	}
	defer f.Close()
	ps, e := parseFile(f)
	if e != nil {
		panic(e)
	}
	out := 0
	for _, p := range ps {
		out += p.value()
	}

	return out, nil
}

func Part2(path string) (int, error) {
	f, e := os.Open(path)
	if e != nil {
		slog.Error(fmt.Sprintf("Cound not open file at %s", path))
		return -1, e
	}
	defer f.Close()
	parseFile(f)
	return -1, nil
}

func parseFile(f *os.File) ([]Pattern, error) {
	buf := bufio.NewScanner(f)
	rowIdx := 0
	rows := make([]int, 0)
	cols := make([]int, 0)
	row := 0
	patterns := make([]Pattern, 0)
	for buf.Scan() {
		line := buf.Text()
		if len(line) <= 1 {
			// New block, reset pattern
			patterns = append(patterns, Pattern{rows, cols})
			rows = make([]int, 0)
			cols = make([]int, 0)
			rowIdx = 0
			continue
		}
		row = 0
		slog.Debug("parsefile", slog.String("line", line))
		for colIdx, ch := range line {
			if len(cols) <= colIdx {
				cols = append(cols, 1)
			}
			if ch == '#' {
				// update row
				if row == 0 {
					row = 1
				}
				row *= PRIMES[colIdx]
				// update col
				cols[colIdx] *= PRIMES[rowIdx]
			}
		}
		rows = append(rows, row)
		rowIdx += 1
	}
	patterns = append(patterns, Pattern{rows, cols})
	for _, p := range patterns {
		slog.Info("parse", slog.Any("p", p))
	}
	return patterns, nil
}

type Pattern struct {
	rows []int
	cols []int
}

func (p *Pattern) value() int {
	return (100*findMirrorIdx(p.rows) +
		findMirrorIdx(p.cols))
}

var PRIMES = []int{
	2,
	3,
	5,
	7,
	11,
	13,
	17,
	19,
	23,
	29,
	31,
	37,
	41,
}

func findMirrorIdx(xs []int) int {
	for i := 0; i < len(xs); i++ {
		isMirrorIdx := true
	inner:
		for o := 0; o <= min(i, len(xs)-i); o++ {
			if xs[i-o] != xs[i+o+1] {
				isMirrorIdx = false
				break inner
			}
		}
		if isMirrorIdx {
			return i
		}
	}
	return -1
}
