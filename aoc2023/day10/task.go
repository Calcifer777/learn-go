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
	loopCells := findLoopCells(&grid, start)
	out := len(loopCells)/2 + len(loopCells)%2
	return out, nil
}

func Part2(path string) (int, error) {
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
	loopCells := findLoopCells(&grid, start)
	_, inside := bisect(grid, loopCells)
	return len(inside), nil
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

func findLoopCells(g *Grid, start Coord) []Coord {
	r, c, prevNS, prevEW := findStartSuccessor(g, start)
	var cell rune
	dir := N
	loopCells := make([]Coord, 0)
	loopCells = append(loopCells, Coord{r: start.r, c: start.c})
loop:
	for {
		cell = (*g)[r][c]
		loopCells = append(loopCells, Coord{r: r, c: c})
		slog.Info("loopTraverse",
			slog.String("cell", string(cell)),
		)
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
		}
		if r == start.r && c == start.c {
			break loop
		}
	}
	slog.Info("startSuccessor",
		slog.Int("loop len", len(loopCells)),
	)
	return loopCells
}

type Direction int

func (d *Direction) String() string {
	return [4]string{"N", "S", "E", "W"}[*d]
}

const (
	N Direction = iota
	S
	E
	W
)

func findStartSuccessor(g *Grid, start Coord) (int, int, *Direction, *Direction) {
	var ew, ns *Direction
	var r, c int
	var up, down, left, right *rune
	var next *rune
	var nextDir, prevDir Direction
	if start.r > 0 {
		up = &(*g)[start.r-1][start.c]
	}
	if start.r < len(*g)-1 {
		down = &(*g)[start.r+1][start.c]
	}
	if start.c > 0 {
		left = &(*g)[start.r][start.c-1]
	}
	if start.c < len(*g)-1 {
		right = &(*g)[start.r][start.c+1]
	}
	if up != nil && (*up == 'F' || *up == '7' || *up == '|') {
		next = up
		r, c = start.r-1, start.c
		dir := S
		ns = &dir
		nextDir = N
	}
	if down != nil && (*down == 'J' || *down == 'L' || *down == '|') {
		if next == nil {
			next = down
			r, c = start.r+1, start.c
			dir := N
			ns = &dir
			nextDir = S
		} else {
			prevDir = S
		}
	}
	if left != nil && (*left == '-' || *left == 'L' || *left == 'F') {
		if next == nil {
			next = left
			r, c = start.r, start.c-1
			dir := E
			ew = &dir
			nextDir = W
		} else {
			prevDir = W
		}
	} else if right != nil && (*right == '-' || *right == 'J' || *right == '7') {
		if next == nil {
			next = right
			r, c = start.r, start.c+1
			dir := W
			ew = &dir
			nextDir = E
		} else {
			prevDir = E
		}
	}
	var replacement rune
	if nextDir == N && prevDir == S {
		replacement = '|'
	} else if nextDir == N && prevDir == W {
		replacement = 'J'
	} else if nextDir == N && prevDir == E {
		replacement = 'L'
	} else if nextDir == S && prevDir == W {
		replacement = '7'
	} else if nextDir == S && prevDir == E {
		replacement = 'F'
	} else if nextDir == W && prevDir == E {
		replacement = '-'
	}
	(*g)[start.r][start.c] = replacement
	slog.Info("startSuccessor",
		slog.Int("row", r),
		slog.Int("col", c),
		slog.Any("ns", ns),
		slog.Any("ew", ew),
		slog.String("prevDir", prevDir.String()),
		slog.String("nextDir", nextDir.String()),
		slog.String("Start repl", string(replacement)),
	)
	return r, c, ns, ew
}

func cellsToMap(cells []Coord) map[Coord]bool {
	m := make(map[Coord]bool)
	for _, c := range cells {
		m[c] = true
	}
	return m
}

func bisect(g Grid, loopCells []Coord) ([]Coord, []Coord) {
	outsideCells := make([]Coord, 0)
	insideCells := make([]Coord, 0)
	loopCellsMap := cellsToMap(loopCells)
	// find cell outside loop
	var out bool
	for r := 0; r < len(g); r++ {
		out = true
		for c := 0; c < len(g[0]); c++ {
			current := Coord{r, c}
			v := g[r][c]
			_, ok := loopCellsMap[current]
			if !ok {
				if out {
					outsideCells = append(outsideCells, current)
				} else {
					insideCells = append(insideCells, current)
				}
			} else {
				if v == '|' || v == 'F' || v == '7' {
					out = !out
				}
			}
			slog.Info("bisect",
				slog.Any("c", current),
				slog.String("v", string(v)),
				slog.Bool("out?", out),
				slog.Bool("in loop?", ok),
				slog.Int("# in", len(insideCells)),
			)
		}
		slog.Info("bisect",
			slog.Any("Inside", len(insideCells)),
			slog.Any("Outside", len(outsideCells)),
		)

	}
	if len(loopCells)+len(insideCells)+len(outsideCells) != len(g[0])*len(g) {
		panic(fmt.Errorf("Error!"))
	}
	return outsideCells, insideCells
}
