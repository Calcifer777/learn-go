package day8

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"regexp"
)

func Part1(path string) (int, error) {
	f, e := os.Open(path)
	if e != nil {
		slog.Error(fmt.Sprintf("Cound not open file at %s", path))
		return -1, e
	}
	defer f.Close()
	net, dirs, e := parseFile(f)
	start := net.nodes[START]
	return traverse(start, dirs), nil
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

func parseFile(f *os.File) (*Net, []Direction, error) {
	buf := bufio.NewScanner(f)
	// read directions
	buf.Scan()
	line := buf.Text()
	dirs := make([]Direction, len(line))
	for i, r := range line {
		if r == 'L' {
			dirs[i] = L
		} else if r == 'R' {
			dirs[i] = R
		}
	}
	slog.Info("parsefile", slog.String("directions", string(dirs)))
	// parse nodes
	re := regexp.MustCompile(`^([(A-Z)]{3}) = \(([(A-Z)]{3}), ([(A-Z)]{3})\)$`)
	nodes := make(map[string]*Node)
	for buf.Scan() {
		line := buf.Text()
		if len(line) == 0 {
			continue
		}
		slog.Info("parsefile", slog.String("line", line))
		names := re.FindStringSubmatch(line)[1:]
		if len(names) != 3 {
			slog.Error("parsefile", slog.Any("names", names))
			return nil, nil, fmt.Errorf("Can't parse line: %s", line)
		}
		for _, n := range names {
			if _, ok := nodes[n]; !ok {
				nodes[n] = NewNode(n, nil, nil)
			}
		}
		if n, ok := nodes[names[0]]; ok {
			n.l = nodes[names[1]]
			n.r = nodes[names[2]]
		}
	}
	for _, node := range nodes {
		slog.Info("parsefile",
			slog.String("node", node.String()),
		)
	}
	return NewNet(nodes), dirs, nil
}

type Direction rune

const (
	L Direction = 'L'
	R Direction = 'R'
)

const (
	START  string = "AAA"
	TARGET string = "ZZZ"
)

type Node struct {
	v    string
	l, r *Node
}

func NewNode(v string, l *Node, r *Node) *Node {
	return &Node{v, l, r}
}

func (n *Node) String() string {
	var lv, rv string
	if n.l == nil {
		lv = ""
	} else {
		lv = n.l.v
	}
	if n.r == nil {
		rv = ""
	} else {
		rv = n.r.v
	}
	return fmt.Sprintf("Node(v=%s, l=%s, r=%s)", n.v, lv, rv)
}

type Net struct {
	nodes map[string]*Node
}

func NewNet(nodes map[string]*Node) *Net {
	return &Net{nodes: nodes}
}

func traverse(start *Node, dirs []Direction) int {
	return innerTraverse(start, dirs, 0)
}

func innerTraverse(n *Node, dirs []Direction, cnt int) int {
	if n.v == TARGET {
		return cnt
	} else {
		var nextNode *Node
		if dirs[cnt%len(dirs)] == L {
			nextNode = n.l
		} else {
			nextNode = n.r
		}
		slog.Info("traverse", slog.String("next", nextNode.v))
		return innerTraverse(nextNode, dirs, cnt+1)
	}
}
