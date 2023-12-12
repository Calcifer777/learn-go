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
		setups += find3(r)
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
	rs, e := parseFile(f)
	if e != nil {
		panic(e)
	}
	setups := 0
	for _, r := range rs {
		rr := repeat(r, 5)
		setups += find(rr)
	}
	return setups, nil
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
		"R(s:`%s`, gs: %v)", //, p: `%s`)",
		r.s,
		r.groups,
		// r.pattern.String(),
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

func repeat(r Record, i int) Record {
	chunks := make([]string, i)
	newGroups := make([]int, 0)
	for j := 0; j <= i; j++ {
		newGroups = append(newGroups, r.groups...)
		chunks = append(chunks, r.s)
	}
	newS := strings.Join(
		chunks,
		"?",
	)
	return NewRecord(newS, newGroups)
}

func find2(r Record) int {
	var looper func(s string, gs []int, lookFor string) int
	looper = func(s string, gs []int, lookFor string) int {
		if len(s) == 0 {
			if len(gs) == 0 {
				return 1
			} else {
				return 0
			}
		} else {
			switch s[0] {
			case '.':
				if lookFor != "*" && lookFor != "." {
					return 0
				} else {
					return looper(s[1:], gs, "*")
				}
			case '#':
				{
					if lookFor != "*" && lookFor != "#" {
						return 0
					} else if len(gs) == 0 {
						return 0
					}
					if gs[0] == 1 {
						return looper(s[1:], gs[1:], ".")
					} else {
						gsNew := append([]int{gs[0] - 1}, gs[1:]...)
						return looper(s[1:], gsNew, "#")
					}
				}
			case '?':
				{
					s1 := "#" + s[1:]
					s2 := "." + s[1:]
					return looper(s1, gs, lookFor) + looper(s2, gs, lookFor)
				}
			default:
				return -1000
			}
		}
	}
	setups := looper(r.s, r.groups, "*")
	slog.Info("find",
		slog.String("s", r.s),
		slog.Int("total", setups),
	)
	return setups
}

func find3(r Record) int {
	patternPrefix := regexp.MustCompile(`^[\.]+`)
	var looper func(s string, gs []int) int
	looper = func(s string, gs []int) int {
		tail := patternPrefix.ReplaceAllLiteralString(s, "")
		if len(tail) == 0 {
			if len(gs) > 0 {
				return 0
			} else {
				return 1
			}
		}
		if len(gs) == 0 {
			if strings.Contains(tail, "#") {
				return 0
			} else {
				return 1
			}
		}
		if tail[0] == '?' {
			return looper("#"+tail[1:], gs) + looper("."+tail[1:], gs)
		}
		re := regexp.MustCompile(fmt.Sprintf(`^[#\?]{%d}[?\.]`, gs[0]))
		match := re.FindString(tail)
		if l := len(match); l > 0 {
			return looper(s[l:], gs[1:])
		} else {
			return 0
		}
	}
	setups := looper(r.s+".", r.groups)
	slog.Info("find",
		slog.String("r", r.String()),
		slog.Int("total", setups),
	)
	return setups
}
