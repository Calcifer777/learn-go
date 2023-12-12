package day12

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"regexp"
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
	rs, e := parseFile(f)
	if e != nil {
		panic(e)
	}
	setups := 0
	for _, r := range rs {
		setups += find(r)
	}

	return setups, nil
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

func parseFile(f *os.File) ([]Record, error) {
	buf := bufio.NewScanner(f)
	records := make([]Record, 0)
	for buf.Scan() {
		line := buf.Text()
		chunks := strings.Fields(line)
		xs := chunks[0]
		groups := make([]int, 0)
		for _, s := range strings.Split(chunks[1], ",") {
			i, _ := strconv.Atoi(s)
			groups = append(groups, i)
		}
		r := NewRecord(xs, groups)
		slog.Info("parsefile",
			slog.String("record", r.String()),
		)
		records = append(records, r)
	}
	return records, nil
}

type Record struct {
	s       string
	groups  []int
	pattern *regexp.Regexp
}

func NewRecord(xs string, gs []int) Record {
	patternStr := `^\.*`
	for i := 0; i < len(gs); i++ {
		patternStr += fmt.Sprintf(`[#]{%d}`, gs[i])
		if i < len(gs)-1 {
			patternStr += `\.+`
		} else {
			patternStr += `\.*$`
		}
	}
	pattern := regexp.MustCompile(patternStr)
	return Record{s: xs, groups: gs, pattern: pattern}
}

func (r *Record) String() string {
	return fmt.Sprintf(
		"R(s:`%s`, gs: %v, p: `%s`)",
		r.s,
		r.groups,
		r.pattern.String(),
	)
}

func find(r Record) int {
	var looper func(s string, d int) int
	looper = func(s string, d int) int {
		for idx, ch := range s {
			if ch == '?' {
				s1 := s[:idx] + "#" + s[idx+1:]
				s2 := s[:idx] + "." + s[idx+1:]
				return looper(s1, d+1) + looper(s2, d+1)
			}
		}
		var out int
		match := r.pattern.MatchString(s)
		if match {
			out = 1
		} else {
			out = 0
		}
		if match {
			slog.Info("find",
				slog.String("s", s),
				slog.Any("gs", r.groups),
				slog.Int("d", d),
				slog.Bool("m", match),
			)
		}
		return out
	}
	setups := looper(r.s, 0)
	slog.Info("find",
		slog.String("s", r.s),
		slog.Int("total", setups),
	)
	return setups
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
