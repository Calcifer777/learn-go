package day10

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
	grid, start, e := parseFile(f)
	if e != nil {
		panic(e)
	}
	l := loopLen(grid, start)
	var out int
	if l%2 == 1 {
		out = l/2 + 1
	} else {
		out = l / 2
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

type Grid = [][]rune
type Coord struct {
	r, c int
}

func parseFile(f *os.File) (Grid, Coord, error) {
	buf := bufio.NewScanner(f)
	grid := make([][]rune, 0)
	start := Coord{r: 0, c: 0}
	for buf.Scan() {
		line := buf.Text()
		slog.Debug("parsefile", slog.String("line", line))
		row := make([]rune, len(line))
		for col, ch := range line {
			row[col] = ch
			if ch == 'S' {
				start.r = len(grid)
				start.c = col
			}
		}
		grid = append(grid, row)
	}
	slog.Info("parseFile",
		slog.Any("start", start),
	)
	return grid, start, nil
}

func loopLen(g Grid, start Coord) int {
	l := 0
	r, c, prevNS, prevEW := findStartSuccessor(g, start)
	var cell rune
	dir := N
loop:
	for {
		cell = g[r][c]
		slog.Info("loopTraverse",
			slog.String("cell", string(cell)),
		)
		l += 1
		switch cell {
		case '|':
			if *prevNS == S {
				r -= 1
			} else {
				r += 1
			}
		case '-':
			if *prevEW == E {
				c -= 1
			} else {
				c += 1
			}
		case 'L':
			if prevNS != nil && *prevNS == N {
				c += 1
				dir = W
				prevEW = &dir
				prevNS = nil
			} else {
				r -= 1
				dir = S
				prevNS = &dir
				prevEW = nil
			}
		case 'J':
			if prevNS != nil && *prevNS == N {
				c -= 1
				dir = E
				prevEW = &dir
				prevNS = nil
			} else {
				r -= 1
				dir = S
				prevNS = &dir
				prevEW = nil
			}
		case '7':
			if prevEW != nil && *prevEW == W {
				r += 1
				dir = N
				prevNS = &dir
				prevEW = nil
			} else {
				c -= 1
				dir = E
				prevEW = &dir
				prevNS = nil
			}
		case 'F':
			if prevNS != nil && *prevNS == S {
				c += 1
				dir = W
				prevEW = &dir
				prevNS = nil
			} else {
				r += 1
				dir = N
				prevNS = &dir
				prevEW = nil
			}
		case 'S':
			break loop
		}
	}
	slog.Info("startSuccessor",
		slog.Int("loopLen", l),
	)
	return l
}

type Direction int

const (
	N Direction = iota
	S
	E
	W
)

func findStartSuccessor(g Grid, start Coord) (int, int, *Direction, *Direction) {
	var ew, ns *Direction
	var r, c int
	var up, down, left, right *rune
	if start.r > 0 {
		up = &g[start.r-1][start.c]
	}
	if start.r < len(g[0])-1 {
		down = &g[start.r+1][start.c]
	}
	if start.c > 0 {
		left = &g[start.r][start.c-1]
	}
	if start.c < len(g)-1 {
		right = &g[start.r][start.c+1]
	}
	if up != nil && (*up == 'F' || *up == '7' || *up == '|') {
		r, c = start.r-1, start.c
		dir := S
		ns = &dir
	} else if down != nil && (*down == 'J' || *down == 'L' || *down == '|') {
		r, c = start.r+1, start.c
		dir := N
		ns = &dir
	} else if left != nil && (*left == '-' || *left == 'L' || *left == 'F') {
		r, c = start.r, start.c-1
		dir := E
		ew = &dir
	} else if right != nil && (*right == '-' || *right == 'J' || *right == '7') {
		r, c = start.r, start.c+1
		dir := W
		ew = &dir
	}
	slog.Info("startSuccessor",
		slog.Int("row", r),
		slog.Int("col", c),
		slog.Any("ns", ns),
		slog.Any("ew", ew),
		slog.String("value", string(g[r][c])),
	)
	return r, c, ns, ew
}
