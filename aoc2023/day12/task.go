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
	rs, e := parseFile(f)
	if e != nil {
		panic(e)
	}
	setups := 0
	for _, r := range rs {
		rr := repeat(r, 5)
		fmt.Printf("%s\n", rr.String())
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
		groups := make([]uint8, 0)
		for _, s := range strings.Split(chunks[1], ",") {
			i, _ := strconv.ParseUint(s, 10, 8)
			groups = append(groups, uint8(i))
		}
		r := NewRecord(xs, groups)
		slog.Debug("parsefile",
			slog.String("record", r.String()),
		)
		records = append(records, r)
	}
	return records, nil
}

type Record struct {
	s      string
	groups []uint8
}

func NewRecord(xs string, gs []uint8) Record {
	patternStr := `^\.*`
	for i := 0; i < len(gs); i++ {
		patternStr += fmt.Sprintf(`[#]{%d}`, gs[i])
		if i < len(gs)-1 {
			patternStr += `\.+`
		} else {
			patternStr += `\.*$`
		}
	}
	return Record{s: xs, groups: gs}
}

func (r *Record) String() string {
	return fmt.Sprintf(
		"R(s:`%s`, gs: %v)",
		r.s,
		r.groups,
	)
}

func repeat(r Record, i int) Record {
	chunks := make([]string, i)
	newGroups := make([]uint8, 0)
	for j := 0; j < i; j++ {
		newGroups = append(newGroups, r.groups...)
		chunks[j] = r.s
	}
	newS := strings.Join(
		chunks,
		"?",
	)
	return NewRecord(newS, newGroups)
}

func find(r Record) int {
	patternPrefix := regexp.MustCompile(`^[\.]+`)
	cache := NewCache()
	var looper func(s string, gs []uint8) int
	looper = func(s string, gs []uint8) int {
		var out int
		if v, ok := cache.get(s, gs); ok {
			return v
		}
		prefix := patternPrefix.FindString(s)
		tail := s[len(prefix):]
		if len(tail) == 0 {
			if len(gs) > 0 {
				out = 0
			} else {
				out = 1
			}
			cache.set(prefix+s, gs, out)
			return out
		}
		if len(gs) == 0 {
			if strings.Contains(tail, "#") {
				out = 0
			} else {
				out = 1
			}
			cache.set(prefix+s, gs, out)
			return out
		}
		if tail[0] == '?' {
			out = looper("."+tail[1:], gs) + looper("#"+tail[1:], gs)
			cache.set(prefix+s, gs, out)
			return out
		}
		re := regexp.MustCompile(fmt.Sprintf(`^[#\?]{%d}(\?|\.|$)`, gs[0]))
		match := re.FindString(tail)
		if l := len(match); l > 0 {
			out = looper(tail[l:], gs[1:])
			cache.set(prefix+s, gs, out)
			return out
		} else {
			return 0
		}
	}
	setups := looper(r.s, r.groups)
	slog.Debug("f3",
		slog.String("r", r.String()),
		slog.Int("total", setups),
	)
	return setups
}

type State struct {
	s  string
	gs string
}

type Cache struct {
	cache map[State]int
}

func NewCache() Cache {
	c := make(map[State]int)
	return Cache{c}
}

func (c Cache) set(s string, groups []uint8, v int) {
	c.cache[State{s, string(groups)}] = v
}

func (c Cache) get(s string, groups []uint8) (int, bool) {
	v, ok := c.cache[State{s, string(groups)}]
	return v, ok
}
