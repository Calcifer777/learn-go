package day11

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
	gs, e := parseFile(f)
	if e != nil {
		panic(e)
	}
	fullRows, fullCols := getFull(gs)
	d := allDists(gs, fullRows, fullCols)
	return d, nil
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

func parseFile(f *os.File) (Galaxies, error) {
	buf := bufio.NewScanner(f)
	rowIdx := 0
	gs := make(Galaxies, 0)
	for buf.Scan() {
		line := buf.Text()
		slog.Debug("parsefile", slog.String("line", line))
		for colIdx, ch := range line {
			if ch != '.' {
				g := Galaxy{rowIdx, colIdx}
				gs = append(gs, g)
				slog.Info("parsefile", slog.String("line", g.String()))
			}
		}
		rowIdx += 1
	}
	return gs, nil
}

type Galaxy struct {
	r, c int
}

func (g *Galaxy) String() string {
	return fmt.Sprintf("G(r: %d, c: %d)", g.r, g.c)
}

func (this *Galaxy) dist(that *Galaxy) int {
	dr := this.r - that.r
	if dr < 0 {
		dr = dr * -1
	}
	dc := this.c - that.c
	if dc < 0 {
		dc = dc * -1
	}
	return dr + dc
}

type Galaxies []Galaxy

func getFull(gs []Galaxy) ([]int, []int) {
	maxC, maxR := 0, 0
	for _, g := range gs {
		if g.r > maxR {
			maxR = g.r
		}
		if g.c > maxC {
			maxC = g.c
		}
	}
	fullRows := make([]int, maxR+1)
	fullCols := make([]int, maxC+1)
	for _, g := range gs {
		fullRows[g.r] = 1
		fullCols[g.c] = 1
	}
	slog.Info("getFull",
		slog.Any("Rows", fullRows),
		slog.Any("Cols", fullCols),
	)
	return fullRows, fullCols
}

func allDists(gs Galaxies, fullRows []int, fullCols []int) int {
	ds := 0
	for _, gF := range gs {
		for _, gT := range gs {
			d := gF.dist(&gT)
			// adjust for empty
			rF := min(gF.r, gT.r)
			rT := max(gF.r, gT.r)
			cF := min(gF.c, gT.c)
			cT := max(gF.c, gT.c)
			adjR := (rT - rF) - Sum(fullRows[rF:rT])
			adjC := (cT - cF) - Sum(fullCols[cF:cT])
			dAdj := d + adjR + adjC
			ds += dAdj
			slog.Info("allDists",
				slog.String("F", gF.String()),
				slog.String("T", gT.String()),
				slog.Int("D", d),
				slog.Int("DAdj", dAdj),
			)
		}
	}
	return ds / 2
}

func Sum(arr []int) int {
	r := 0
	for _, i := range arr {
		r += i
	}
	return r
}
