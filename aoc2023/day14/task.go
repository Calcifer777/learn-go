package day14

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
	lines, e := parseFile(f)
	if e != nil {
		panic(e)
	}
	out := value(lines)
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

func parseFile(f *os.File) ([]string, error) {
	buf := bufio.NewScanner(f)
	lines := make([]string, 0)
	for buf.Scan() {
		line := buf.Text()
		slog.Debug("parsefile", slog.String("line", line))
		lines = append(lines, line)
	}
	return lines, nil
}

func value(lines []string) int {
	tracker := make([]int, len(lines[0]))
	for i := 0; i < len(tracker); i++ {
		tracker[i] = len(tracker)
	}
	maxV := len(lines)
	values := make([]int, len(lines[0]))
	for rowIdx, line := range lines {
		for colIdx, ch := range line {
			switch ch {
			case '.':
				continue
			case '#':
				tracker[colIdx] = maxV - rowIdx - 1
			case 'O':
				{
					values[colIdx] += tracker[colIdx]
					tracker[colIdx] -= 1
				}
			}
		}
	}
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}
