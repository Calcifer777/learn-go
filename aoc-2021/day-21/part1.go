package main

import (
	"fmt"
	"encoding/json"
	"go.uber.org/zap"
  "io/ioutil"
)

var sugarLogger *zap.SugaredLogger

func WithDevelopmentLogger() {
  jsonFile, _ := ioutil.ReadFile("log-config.json")
  var config zap.Config
  if err := json.Unmarshal(jsonFile, &config); err != nil {
    panic(err)
  }
	// config := zap.NewDevelopmentConfig()
	logger, _ := config.Build()
	sugarLogger = logger.Sugar()
}

func WithProductionLogger() {
	config := zap.NewProductionConfig()
	level := zap.NewAtomicLevelAt(zap.WarnLevel)
	config.Level = level
	logger, _ := config.Build()
	sugarLogger = logger.Sugar()
}

type Player struct {
	name  string
	turn  int
	pos   int
	score int
}

type Dice interface {
	Roll() int
}

type DDice struct {
	n int
}

func (d *DDice) Roll() int {
	v := (d.n % 100) + 1
	d.n += 1
	return v
}

type Game struct {
	p1   Player
	p2   Player
	dice Dice
	next int
}

func (g *Game) PlayTurn() bool {
	var p *Player
	if g.next == 1 {
		p = &g.p1
	} else {
		p = &g.p2
	}
	v := 0
	for i := 0; i < 3; i++ {
		v += g.dice.Roll()
	}
	p.pos = ((p.pos + v - 1) % 10) + 1
	p.score += p.pos
	sugarLogger.Infof("%s: rolled %d, moving to %d, score: %d", p.name, v, p.pos, p.score)
	p.turn += 1
	g.next = (g.next % 2) + 1
	return p.score >= 1000
}

func (g *Game) Play() int {
	sugarLogger.Infof("Starting game")
	for {
		won := g.PlayTurn()
		if won {
			switch g.next {
			case 1:
				return 2
			case 2:
				return 1
			default:
				panic("Game ended, but unexpected next player")
			}
		}
	}
}

func (g Game) String() string {
	s := "Game details:\n"
	switch v := g.dice.(type) {
	case *DDice:
		s += fmt.Sprintf("\tRolls: %+v\n", v.n)
	default:
		panic("unexpected")
	}
	if g.p1.score >= 1000 {
		s += fmt.Sprintf("\tStatus: %s won\n", g.p1.name)
	} else if g.p2.score >= 1000 {
		s += fmt.Sprintf("\tStatus: %s won\n", g.p2.name)
	} else {
		s += fmt.Sprintf("\tStatus: in progress\n")
	}
	s += fmt.Sprintf("\t%s score: %d\n", g.p1.name, g.p1.score)
	s += fmt.Sprintf("\t%s score: %d", g.p2.name, g.p2.score)
	return s
}

func main() {
	WithDevelopmentLogger()
	// WithProductionLogger()
	game := Game{
		p1:   Player{name: "P1", turn: 0, pos: 2, score: 0},
		p2:   Player{name: "P2", turn: 0, pos: 7, score: 0},
		dice: &DDice{n: 0},
		next: 1,
	}
	game.Play()
	fmt.Printf("%s\n", game)
	var losingPlayer *Player
	if game.p2.score >= 1000 {
		losingPlayer = &game.p1
	} else {
		losingPlayer = &game.p2
	}
	switch v := game.dice.(type) {
	case *DDice:
		fmt.Printf("Part1 -> %d\n", losingPlayer.score*v.n)
	default:
		panic("unexpected")
	}
}
