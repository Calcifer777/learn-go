package day2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Draw = map[string]int
type Game = []Draw
type Games = []Game
type Bag = map[string]int

func parseLine(line string) (Game, error) {
	chunks := strings.SplitN(line, ":", 2)
	// parse game id
	head := chunks[0]
	gameID := strings.Fields(head)[1]
	_, err := strconv.Atoi(gameID)
	if err != nil {
		return nil, err
	}
	// parse draws
	tail := chunks[1]
	game := make(Game, 0)
	for _, chunk := range strings.Split(tail, ";") {
		draw := make(Draw, 0)
		for _, color := range strings.Split(chunk, ",") {
			xs := strings.Fields(color)
			n, err := strconv.Atoi(xs[0])
			if err != nil {
				return nil, err
			}
			draw[xs[1]] = n
		}
		game = append(game, draw)
	}
	return game, nil
}

func parseFile(f *os.File) []Game {
	scanner := bufio.NewScanner(f)
	games := make([]Game, 0)
	for scanner.Scan() {
		draws, err := parseLine(scanner.Text())
		if err != nil {
			panic(err)
		}
		games = append(games, draws)
	}
	return games

}
