package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"utils"
)

type RiskMap [][]int
type Coord struct{ r, c int }
type Node struct {
	coord Coord
	risk  int
	prev  *Node
}

func (n *Node) Path() []Coord {
	path := []Coord{n.coord}
	node := n
	for {
		if node.prev == nil {
			break
		} else {
			path = append(path, node.prev.coord)
			node = node.prev
		}
	}
	return path
}

type Queue []Node

func ParseInput(lines []string) RiskMap {
	riskMap := make(RiskMap, len(lines))
	for r, line := range lines {
		row := make([]int, len(lines[0]))
		for c, char := range strings.Split(line, "") {
			i, _ := strconv.Atoi(char)
			row[c] = i
		}
		riskMap[r] = row
	}
	return riskMap
}

func Dijkstra(riskMap RiskMap) (Node, error) {
	height := len(riskMap)
	width := len(riskMap[0])
	start := Coord{0, 0}
	target := Coord{r: len(riskMap) - 1, c: len(riskMap[0]) - 1}
	// Initialize queue from starting node (0, 0)
	queue := []Node{Node{coord: start, risk: 0, prev: nil}}
	// Initialize bestRisks map; it keeps track of the risks of expanded nodes
	// and it is used to expand nodes only from the best path
	// It prevents from visiting the same Node twice in the same path, or
	// from enqueueing a Node which was already visited from a less costly path
	bestRisks := make(map[Coord]int)
	for {
		if len(queue) == 0 {
			return Node{}, errors.New("Exhausted available nodes, could not find a path to the target coord")
		}
		if queue[0].coord == target {
			return queue[0], nil
		}
		head := queue[0]
		queue = queue[1:]
		candidates := [4]Coord{
			Coord{r: head.coord.r + 1, c: head.coord.c + 0}, // down
			Coord{r: head.coord.r - 1, c: head.coord.c + 0}, // up
			Coord{r: head.coord.r + 0, c: head.coord.c + 1}, // right
			Coord{r: head.coord.r + 0, c: head.coord.c - 1}, // left
		}
		for _, candidate := range candidates {
			if candidate.r >= 0 && candidate.r < height && // check vertical bounds
				candidate.c >= 0 && candidate.c < width { // check horizontal bounds
				currentNode := Node{
					coord: candidate,
					risk:  head.risk + riskMap[candidate.r][candidate.c],
					prev:  &head,
				}
				// visit a Node only from the best path
				if br, ok := bestRisks[candidate]; !ok || currentNode.risk < br {
					queue = append(queue, currentNode)
					bestRisks[candidate] = currentNode.risk
				}
			}
		}
		// Reorder queue by risk
		sort.Slice(queue, func(i, j int) bool { return queue[i].risk < queue[j].risk })
	}
	return Node{coord: start, risk: 0, prev: nil}, errors.New("Could not find a path to the target coord")
}

// Replicate a RiskMap n times vertically and horizontally,
// increasing the risk values at each increment
// (diagonal increment counts as 2)
// Risk values wrap at 1 when over 10 (e.g. 3 + 1 -> 4, 9 + 1 -> 1)
func Extend(m RiskMap, n int) RiskMap {
	height := len(m)
	width := len(m[0])
	newMap := make(RiskMap, height*n)
	for r := 0; r < height*n; r++ {
		row := make([]int, width*n)
		for c := 0; c < width*n; c++ {
			row[c] = (m[r%height][c%width]+r/height+c/width-1)%9 + 1
		}
		newMap[r] = row
	}
	return newMap
}

func Part1() {
	lines, _ := utils.ReadLines("input.txt")
	riskMap := ParseInput(lines)
	node, err := Dijkstra(riskMap)
	utils.Check(err)
	fmt.Printf("Part 1 -> %d\n", node.risk)
}

func Part2() {
	lines, _ := utils.ReadLines("input.txt")
	riskMap := ParseInput(lines)
	extended := Extend(riskMap, 5)
	node, err := Dijkstra(extended)
	utils.Check(err)
	fmt.Printf("Part 2 -> %d\n", node.risk)
}

func main() {
	Part1()
	Part2()
}
