package day5

import (
	"bufio"
	"fmt"
	"log/slog"
	"math"
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
	seeds, convMaps, e := parseFile(f)
	if e != nil {
		slog.Error(fmt.Sprintf("Cound not parse file"))
		return -1, e
	}
	locations := make([]int, len(seeds))
	for idx, seed := range seeds {
		loc := traverseMappings(seed, convMaps)
		locations[idx] = loc
	}
	minLocation := min(locations)

	return minLocation, nil
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

type Range struct {
	dst, src, rng int
}

func NewRange(dst, src, rng int) Range {
	return Range{dst, src, rng}
}

type ConvMap struct {
	name   string
	ranges []Range
}

func NewConvMap(name string, ranges []Range) ConvMap {
	return ConvMap{name, ranges}
}

func parseFile(f *os.File) ([]int, []ConvMap, error) {
	buf := bufio.NewScanner(f)
	// parse seeds
	buf.Scan()
	firstLine := buf.Text()
	seeds := make([]int, 0)
	for _, s := range strings.Fields(strings.Split(firstLine, ": ")[1]) {
		seed, e := strconv.Atoi(s)
		if e != nil {
			slog.Debug("Could not parse seeds")
			return nil, nil, e
		} else {
			seeds = append(seeds, seed)
		}
	}
	// parse maps
	convMaps := make([]ConvMap, 0)
	ranges := make([]Range, 0)
	mapId := ""
	for buf.Scan() {
		line := buf.Text()
		if len(line) == 0 {
			if len(ranges) > 0 {
				convMaps = append(convMaps, NewConvMap(mapId, ranges))
			}
			ranges = make([]Range, 0)
			mapId = ""
		} else if len(ranges) == 0 && mapId == "" {
			mapId = strings.Fields(line)[0]
		} else {
			range_ := make([]int, 3)
			for idx, s := range strings.Fields(line) {
				n, e := strconv.Atoi(s)
				if e != nil {
					slog.Debug("Error during maps parsing")
					return nil, nil, e
				}
				range_[idx] = n
			}
			ranges = append(ranges, NewRange(range_[0], range_[1], range_[2]))
		}
	}
	// add last map
	if len(ranges) > 0 && mapId != "" {
		convMaps = append(convMaps, NewConvMap(mapId, ranges))
	}

	slog.Debug("parseFileOutput",
		slog.Any("seeds", seeds),
		slog.Any("convMaps", convMaps),
	)
	return seeds, convMaps, nil
}

func traverseMappings(seed int, convMaps []ConvMap) int {
	dst := seed
outer:
	for _, c := range convMaps {
		slog.Info("Traverse step",
			slog.Int("v", dst),
		)
		for _, r := range c.ranges {
			offset := dst - r.src
			if offset >= 0 && offset <= r.rng {
				dst = r.dst + offset
				continue outer
			}
		}
	}
	slog.Debug("traverse result",
		slog.Int("dst", dst),
	)
	return dst
}

func min(arr []int) int {
	minValue := math.MaxInt64
	for _, x := range arr {
		if x < minValue {
			minValue = x
		}
	}
	slog.Debug(
		"MinLocation",
		slog.Int("value", minValue),
	)
	return minValue
}
