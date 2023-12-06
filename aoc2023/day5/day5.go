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
	minLocation := arrMin(locations)

	return minLocation, nil
}

func Part2(path string) (int, error) {
	f, e := os.Open(path)
	if e != nil {
		slog.Error(fmt.Sprintf("Cound not open file at %s", path))
		return -1, e
	}
	defer f.Close()
	seeds, convMaps, e := parseFile2(f)
	if e != nil {
		panic(e)
	}
	minValue := math.MaxInt64
	for _, s := range seeds {
		mappedS := mapRangeIter(s, convMaps)
		minStart := findMinSeedRange(mappedS)
		if minStart < minValue {
			minValue = minStart
		}
	}
	return minValue, nil
}

type Range struct {
	dst_start, src_start, span int
}

func NewRange(dst_start, src_start, span int) Range {
	return Range{dst_start, src_start, span}
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
	//
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
			offset := dst - r.src_start
			if offset >= 0 && offset <= r.span {
				dst = r.dst_start + offset
				continue outer
			}
		}
	}
	slog.Debug("traverse result",
		slog.Int("dst", dst),
	)
	return dst
}

func arrMin(arr []int) int {
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

func parseFile2(f *os.File) ([]SeedRange, []ConvMap, error) {
	buf := bufio.NewScanner(f)
	// parse seeds
	buf.Scan()
	firstLine := buf.Text()
	seedFields := strings.Fields(strings.Split(firstLine, ": ")[1])
	seeds := make([]SeedRange, 0)
	for i := 0; i < len(seedFields); i += 2 {
		seedRangeStart, e1 := strconv.Atoi(seedFields[i])
		seedRangeEnd, e2 := strconv.Atoi(seedFields[i+1])
		if e1 != nil {
			slog.Debug("Could not parse seeds")
			return nil, nil, e1
		}
		if e2 != nil {
			slog.Debug("Could not parse seeds")
			return nil, nil, e2
		}
		seeds = append(seeds, *newSeed(seedRangeStart, seedRangeEnd))
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
	//
	for _, s := range seeds {
		slog.Debug("parseFileOutput", slog.Any("seed", s))
	}
	for _, m := range convMaps {
		slog.Debug("parseFileOutput", slog.Any("map", m))
	}
	return seeds, convMaps, nil
}

type SeedRange struct {
	start, span int
}

func newSeed(start, span int) *SeedRange {
	return &SeedRange{start, span}
}

func overlap(start1, span1, start2, span2 int) (int, int, bool) {
	minOverlap := max(start1, start2)
	maxOverlap := min(start1+span1-1, start2+span2-1)
	ok := true
	if minOverlap > maxOverlap {
		minOverlap = -1
		maxOverlap = -1
		ok = false
	}
	slog.Debug("Overlap",
		slog.Int("start1", start1),
		slog.Int("span1", span1),
		slog.Int("start2", start2),
		slog.Int("span2", span2),
		slog.Bool("ok", ok),
		slog.Int("minOverlap", minOverlap),
		slog.Int("maxOverlap", maxOverlap),
	)
	return minOverlap, maxOverlap, ok
}

func findRanges(s SeedRange, minOverlap, maxOverlap int) (*SeedRange, *SeedRange, *SeedRange) {
	if minOverlap == s.start && maxOverlap == s.start+s.span {
		return nil, &s, nil
	}
	var before, after, toMap *SeedRange
	if minOverlap > s.start {
		before = newSeed(s.start, minOverlap-s.start)
	}
	if maxOverlap < s.start+s.span-1 {
		after = newSeed(maxOverlap+1, s.start+s.span)
	}
	toMapMin := max(s.start, minOverlap)
	toMapMax := min(s.start+s.span, maxOverlap)
	toMap = newSeed(toMapMin, toMapMax-toMapMin+1)
	return before, toMap, after
}

func mapRange(seeds []SeedRange, convMap ConvMap) []SeedRange {
	mapped := make([]SeedRange, 0)
	for _, r := range convMap.ranges {
		acc := make([]SeedRange, 0)
		for _, s := range seeds {
			minOverlap, maxOverlap, ok := overlap(s.start, s.span, r.src_start, r.span)
			if !ok {
				acc = append(acc, s)
			} else {
				before, toMap, after := findRanges(s, minOverlap, maxOverlap)
				slog.Debug("mapRange",
					slog.Any("before", before),
					slog.Any("toMap", toMap),
					slog.Any("after", after),
				)
				if before != nil {
					acc = append(acc, *before)
				}
				if after != nil {
					acc = append(acc, *after)
				}
				offset := r.dst_start - r.src_start
				mapped = append(mapped, *newSeed(toMap.start+offset, toMap.span))
			}
		}
		seeds = acc
	}
	for _, s := range seeds {
		mapped = append(mapped, s)
	}
	// Logs
	for _, s := range mapped {
		slog.Debug("MapRange",
			slog.Any("Mapped range", s),
		)
	}
	return mapped
}

func mapRangeIter(seed SeedRange, convMaps []ConvMap) []SeedRange {
	seeds := make([]SeedRange, 1)
	seeds[0] = seed
	for _, cm := range convMaps {
		seeds = mapRange(seeds, cm)
		break
	}
	return seeds
}

func findMinSeedRange(seeds []SeedRange) int {
	minValue := math.MaxInt64
	for _, s := range seeds {
		if s.start < minValue {
			minValue = s.start
		}
	}
	return minValue
}
