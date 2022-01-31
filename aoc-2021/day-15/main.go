package main

import (
	// "container/heap"
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
	bestRisks := make(map[Coord]int)
	for {
		// fmt.Printf("Step: %d\n", i)
		if len(queue) == 0 {
			return Node{}, errors.New("Exhausted available nodes, could not find a path to the target coord")
		}
		if queue[0].coord == target {
			return queue[0], nil
		}
		head := queue[0]
		queue = queue[1:]
		// fmt.Printf("Queue:\n"); for _, n := range queue { fmt.Printf("%+v\n", n) }
		candidates := [4]Coord{
			Coord{r: head.coord.r + 1, c: head.coord.c + 0}, // down
			Coord{r: head.coord.r - 1, c: head.coord.c + 0}, // up
			Coord{r: head.coord.r + 0, c: head.coord.c + 1}, // right
			Coord{r: head.coord.r + 0, c: head.coord.c - 1}, // left
		}
		for _, candidate := range candidates {
			if candidate.r >= 0 && candidate.r < height && // do not go out of bound horizontally
				candidate.c >= 0 && candidate.c < width { // do not go out of bound vertically
				currentNode := Node{
					coord: candidate,
					risk:  head.risk + riskMap[candidate.r][candidate.c],
					prev:  &head,
				}
				// do not visit same node twice
				if br, ok := bestRisks[candidate]; !ok || currentNode.risk < br {
					fmt.Printf("Adding node to queue: %+v\n", candidate)
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

func main() {
	lines, _ := utils.ReadLines("input.txt")
	// lines, _ := utils.ReadLines("input-sample-1.txt")
	riskMap := ParseInput(lines)
	node, err := Dijkstra(riskMap)
	utils.Check(err)
	fmt.Printf("Reached %v, Risk %d, with path: \n", node.coord, node.risk)
	// fmt.Printf("With path: \n")
	// for _, c := range node.Path() {
	// 	fmt.Printf("%v\n", c)
	// }
}
