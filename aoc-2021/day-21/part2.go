package main

import (
	// "bufio"
	"fmt"
	// "os"
)

const winningScore = 21

// LIFO queue implementation
type Queue struct {
	head *DiracGame
	tail *Queue
}

func (q Queue) Pop() (Queue, *DiracGame) {
	if q.tail != nil {
		return Queue{q.tail.head, q.tail.tail}, q.head
	} else {
		return Queue{nil, nil}, q.head
	}
}

func (q Queue) Push(v DiracGame) Queue {
	return Queue{head: &v, tail: &q}
}

func (q Queue) HasNext() bool {
	return q.head != nil
}

func (q Queue) Size() int {
	size := 0
	queue := q
	for {
		if !queue.HasNext() {
			break
		} else {
			queue, _ = queue.Pop()
			size++
		}
	}
	return size
}

type Player struct {
	name  string
	pos   int
	score int
}

type DiracGame struct {
	p1           Player
	p2           Player
	onTurn       *Player
	turns        int
	currentRolls int
	status       GameStatus
  weight       int64
}

func (g DiracGame) String() string {
	s := "Game:\n"
	if g.p1.score >= winningScore {
		s += fmt.Sprintf("\tStatus: %s won\n", g.p1.name)
	} else if g.p2.score >= winningScore {
		s += fmt.Sprintf("\tStatus: %s won\n", g.p2.name)
	} else {
		s += fmt.Sprintf("\tStatus: in progress\n")
	}
	s += fmt.Sprintf("\tTurns: %d\n", g.turns)
	s += fmt.Sprintf("\t%s at pos %d, score: %d\n", g.p1.name, g.p1.pos, g.p1.score)
	s += fmt.Sprintf("\t%s at pos %d, score: %d\n", g.p2.name, g.p2.pos, g.p2.score)
	s += fmt.Sprintf("\tTurn:\n")
	s += fmt.Sprintf("\t\tPlayer      %s\n", g.onTurn.name)
	return s
}

func (g DiracGame) PlayTurn(v int, weight int64) DiracGame {
	p1 := g.p1
	p2 := g.p2
	status := g.status
	turns := g.turns
  weight = g.weight * weight
  // Select player on the play
	var onTurn *Player
	if g.onTurn.name == p1.name {
		onTurn = &p1
	} else {
		onTurn = &p2
	}
  // Get new position and score
  onTurn.pos = ((onTurn.pos + v - 1) % 10) + 1
  onTurn.score += onTurn.pos
  // Update game status
  if onTurn.score >= winningScore {
    status = Done
  }
  // Update player on the play
  if *onTurn == p1 {
    onTurn = &p2
  } else if *onTurn == p2 {
    onTurn = &p1
  } else {
    panic("Unreachable: can't switch turns")
  }
  turns += 1
  // Return new game status
	newGame := DiracGame{
		p1:           p1,
		p2:           p2,
		status:       status,
		turns:        turns,
		onTurn:       onTurn,
		weight:       weight,
	}
	return newGame
}

func PlayDiracTurn(g DiracGame) []DiracGame {
	newGames := make([]DiracGame, 7)
  rollOutcomes := map[int]int64 {
    3: 1,
    4: 3,
    5: 6,
    6: 7,
    7: 6,
    8: 3,
    9: 1,
  }
  idx := 0
  for rollValue, weight := range rollOutcomes {
		newGames[idx] = g.PlayTurn(rollValue, weight)
    idx++
	}
	return newGames
}

type GameStatus int

const (
	InProgress GameStatus = iota
	Done
)

func main() {
	p1 := Player{name: "Player 1", pos: 4, score: 0}
	p2 := Player{name: "Player 2", pos: 8, score: 0}
	start := DiracGame{
		p1:           p1,
		p2:           p2,
		status:       InProgress,
		turns:        0,
		onTurn:       &p1,
		weight:       1,
	}
	gamesQueue := Queue{&start, nil}
	var game *DiracGame
	var p1Wins, p2Wins int64
	var i int
	// for i := 0; i < 10e7; i++ {
	for {
		if !gamesQueue.HasNext() {
			break
		}
		gamesQueue, game = gamesQueue.Pop()
		// fmt.Printf("Expanding\n%s\n", *game)
		for _, g := range PlayDiracTurn(*game) {
			switch g.status {
			case InProgress:
				{
					// fmt.Printf("Pushing\n%s\n", g)
					gamesQueue = gamesQueue.Push(g)
				}
			case Done:
				{
					if g.p1.score >= winningScore {
						// fmt.Printf("Player 1 wins!\n")
						p1Wins += g.weight
					} else if g.p2.score >= winningScore {
						// fmt.Printf("Player 2 wins with score %d!\n", g.p2.score)
						p2Wins += g.weight
					} else {
						panic("unreachable")
					}
				}
			default:
				panic("unreachable")
			}
		}
		if i%10e5 == 0 {
			fmt.Printf("Queue size: %d\n", gamesQueue.Size())
			fmt.Printf("Player1 wins: %d\n", p1Wins)
			fmt.Printf("Player2 wins: %d\n", p2Wins)
			fmt.Printf("\n")
		}
		i++
	}

}
