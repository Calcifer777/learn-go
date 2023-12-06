package day6

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
	races, e := parseFile(f)
	if e != nil {
		panic(e)
	}
	totalImprovements := 1
	for _, r := range races {
		numImprovements := getImprovements(r)
		slog.Info("Part1",
			slog.Any("Race", r),
			slog.Int("numImprovements", numImprovements),
		)
		totalImprovements *= numImprovements
	}
	return totalImprovements, nil
}

func Part2(path string) (int, error) {
	f, e := os.Open(path)
	if e != nil {
		slog.Error(fmt.Sprintf("Cound not open file at %s", path))
		return -1, e
	}
	defer f.Close()
	race, e := parseFile2(f)
	if e != nil {
		panic(e)
	}
	numImprovements := getImprovements(race)
	return numImprovements, nil
}

func parseFile(f *os.File) ([]Race, error) {
	buf := bufio.NewScanner(f)
	// parse times
	buf.Scan()
	line := buf.Text()
	times := make([]int, 0)
	for _, s := range strings.Fields(strings.Split(line, ":")[1]) {
		i, e := strconv.Atoi(s)
		if e != nil {
			panic(e)
		}
		times = append(times, i)
	}
	// parse distances
	buf.Scan()
	line = buf.Text()
	distances := make([]int, 0)
	for _, s := range strings.Fields(strings.Split(line, ":")[1]) {
		i, e := strconv.Atoi(s)
		if e != nil {
			panic(e)
		}
		distances = append(distances, i)
	}
	if len(times) != len(distances) {
		panic(fmt.Sprintf(
			"Different lengths for times (%d) and distances (%d)",
			len(times),
			len(distances),
		))
	}
	// make races
	races := make([]Race, 0)
	for i := 0; i < len(times); i++ {
		r := newRace(times[i], distances[i])
		slog.Info("parseFile",
			slog.String("Race", r.ToString()),
		)
		races = append(races, r)
	}
	return races, nil
}

type Race struct {
	time       int
	RecordDist int
}

func newRace(time, dist int) Race {
	return Race{time, dist}
}

func (r *Race) ToString() string {
	return fmt.Sprintf("Race(time: %d, dist: %d)", r.time, r.RecordDist)
}

func dist(charge, maxTime int) int {
	return charge * (maxTime - charge)
}

func maxDist(maxTime int) int {
	bestCharge := maxTime / 2
	return bestCharge * (maxTime - bestCharge)
}

// -c ^ 2 + maxTime*c - record
func improvementRange(maxTime int, record int) (int, int) {
	delta := math.Sqrt(
		math.Pow(float64(maxTime), 2) -
			4*float64(record))
	rangeFloor := int(math.Ceil((float64(maxTime) - delta) / 2))
	rangeCeil := int(math.Floor((float64(maxTime) + delta) / 2))
	if dist(rangeFloor, maxTime)-record == 0 {
		rangeFloor += 1
	}
	if dist(rangeCeil, maxTime)-record == 0 {
		rangeCeil -= 1
	}
	return rangeFloor, rangeCeil
}

func getImprovements(r Race) int {
	improvRangeMin, improvRangeMax := improvementRange(r.time, r.RecordDist)
	numImprovements := improvRangeMax - improvRangeMin + 1
	slog.Info("getImprovements",
		slog.String("Race", r.ToString()),
		slog.Int("MaxDist", maxDist(r.time)),
		slog.Any("Better min", improvRangeMin),
		slog.Any("Better max", improvRangeMax),
		slog.Any("Num. possible improvements", numImprovements),
	)
	return numImprovements
}

func parseFile2(f *os.File) (Race, error) {
	buf := bufio.NewScanner(f)
	// parse times
	buf.Scan()
	line := buf.Text()
	timeStr := strings.Replace(
		strings.Split(line, ":")[1], " ", "", -1,
	)
	time, e := strconv.Atoi(timeStr)
	if e != nil {
		panic(e)
	}
	// parse distances
	buf.Scan()
	line = buf.Text()
	distanceStr := strings.Replace(
		strings.Split(line, ":")[1], " ", "", -1,
	)
	distance, e := strconv.Atoi(distanceStr)
	if e != nil {
		panic(e)
	}
	// make races
	race := newRace(time, distance)
	slog.Info("parseFile2",
		slog.String("Race", race.ToString()),
	)
	return race, nil
}
