package day9

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

func Part1(path string) (int, error) {
	f, e := os.Open(path)
	if e != nil {
		slog.Error(fmt.Sprintf("Cound not open file at %s", path))
		return -1, e
	}
	defer f.Close()
	hs, e := parseFile(f)
	hDiffs := FMap(hDiff, hs)
	preds := FMap(getPred, hDiffs)
	out := part1Out(preds)
	return out, nil
}

func Part2(path string) (int, error) {
	f, e := os.Open(path)
	if e != nil {
		slog.Error(fmt.Sprintf("Cound not open file at %s", path))
		return -1, e
	}
	defer f.Close()
	return -1, nil
}

func parseFile(f *os.File) ([]History, error) {
	buf := bufio.NewScanner(f)
	histories := make([][]int, 0)
	for buf.Scan() {
		line := buf.Text()
		slog.Info("parsefile", slog.String("line", line))
		chunks := strings.Fields(line)
		h := make([]int, len(chunks))
		for idx, c := range chunks {
			i, _ := strconv.Atoi(c)
			h[idx] = i
		}
		histories = append(histories, h)
		slog.Info("parsefile", slog.Any("H", h))
	}
	return histories, nil
}

type History = []int

func allZeros(xs []int) bool {
	for _, x := range xs {
		if x != 0 {
			return false
		}
	}
	return true
}

func hDiff(h History) []History {
	ds := make([]History, 0)
	arr := h
	ds = append(ds, arr)
	for len(arr) > 1 && !allZeros(arr) {
		d := make(History, 0)
		for i := 1; i < len(arr); i++ {
			d = append(d, arr[i]-arr[i-1])
		}
		ds = append(ds, d)
		slog.Info("hDiff",
			slog.Any("d", d),
		)
		arr = d
	}
	return ds
}

func getPred(d []History) int {
	out := 0
	for _, h := range d {
		out += h[len(h)-1]
	}
	slog.Info(
		"getPred",
		slog.Any("d", d),
		slog.Int("pred", out),
	)
	return out
}

func FMap[T any, S any](fn func(T) S, xs []T) []S {
	acc := make([]S, len(xs))
	for idx, t := range xs {
		acc[idx] = fn(t)
	}
	return acc
}

func part1Out(preds []int) int {
	out := 0
	for _, p := range preds {
		out += p
	}
	slog.Info(
		"part1Out",
		slog.Int("out", out),
	)
	return out
}
