package day4

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
	games, e := parseFile(f)
	if e != nil {
		slog.Error(fmt.Sprintf("Cound not parse file: %v", e))
		return -1, e
	}
	sum := 0
	for _, g := range games {
		slog.Debug(
			"Game parsing output",
			slog.String("Game", fmt.Sprintf("%v", *g)),
		)
		score := getScore(g)
		slog.Debug(
			"Output",
			slog.Int("Score", score),
		)
		sum += score
	}
	return sum, nil
}

func Part2(path string) (int, error) {
	f, e := os.Open(path)
	if e != nil {
		slog.Error(fmt.Sprintf("Cound not open file at %s", path))
		return -1, e
	}
	defer f.Close()
	games, e := parseFile(f)
	if e != nil {
		slog.Error(fmt.Sprintf("Cound not parse file: %v", e))
		return -1, e
	}
	scratch := make(map[int]int)
	for _, g := range games {
		scratch[g.id] += 1
		slog.Debug(
			"Game parsing output",
			slog.String("Game", fmt.Sprintf("%v", *g)),
		)
		winning := getScorePart2(g)
		for _, w := range winning {
			scratch[w] += scratch[g.id]
		}
		slog.Debug(
			"Output",
			slog.Any("Scratch", scratch),
		)
	}
	sum := 0
	for _, v := range scratch {
		sum += v
	}
	return sum, nil
}

type Game struct {
	id      int
	winning []int
	played  []int
}

func parseFile(f *os.File) ([]*Game, error) {
	buf := bufio.NewScanner(f)
	games := make([]*Game, 0)
	for buf.Scan() {
		line := buf.Text()
		slog.Debug(line)
		game, e := parseLine(line)
		if e != nil {
			slog.Debug(fmt.Sprintf("Could parse game"))
			return nil, e
		}
		games = append(games, game)
	}
	return games, nil
}

func parseLine(line string) (*Game, error) {
	chunks := strings.Split(line, ": ")
	// parse id
	gameIdStr := strings.Fields(chunks[0])[1]
	gameId, e := strconv.Atoi(gameIdStr)
	if e != nil {
		slog.Debug(fmt.Sprintf("Could not get game id!"))
		return nil, e
	}
	// parse numbers
	numberChunks := strings.Split(chunks[1], "|")
	// winning
	winning := make([]int, 0)
	for _, n := range strings.Fields(numberChunks[0]) {
		i, e := strconv.Atoi(n)
		if e != nil {
			slog.Debug(fmt.Sprintf("Could not parse winning numbers"))
			return nil, e
		}
		winning = append(winning, i)
	}
	played := make([]int, 0)
	// played
	for _, n := range strings.Fields(numberChunks[1]) {
		i, e := strconv.Atoi(n)
		if e != nil {
			slog.Debug(fmt.Sprintf("Could not parse played numbers"))
			return nil, e
		}
		played = append(played, i)
	}
	return &Game{gameId, winning, played}, nil
}

func getScore(g *Game) int {
	score := 0
	correct := 0
	for _, n := range g.played {
	inner:
		for _, w := range g.winning {
			if n == w {
				if correct == 0 {
					score = 1
				} else {
					score *= 2
				}
				slog.Debug(fmt.Sprintf("Got number: %d - Score: %d", n, score))
				correct += 1
				break inner
			}
		}
	}
	return score
}

func getScorePart2(g *Game) []int {
	correct := make([]int, 0)
	for _, n := range g.played {
	inner:
		for _, w := range g.winning {
			if n == w {
				toAppend := g.id + 1 + len(correct)
				slog.Debug(fmt.Sprintf("Got scratch: %d", toAppend))
				correct = append(correct, toAppend)
				break inner
			}
		}
	}
	return correct
}
