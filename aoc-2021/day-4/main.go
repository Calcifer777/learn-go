package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

type BoardNumber struct {
	value     int
	extracted bool
}

type Board [][]BoardNumber

func ParseLine(line string) []BoardNumber {
	values := strings.Fields(line)
	numbers := make([]BoardNumber, 5)
	for idx, i := range values {
		v, _ := strconv.Atoi(i)
		numbers[idx] = BoardNumber{v, false}
	}
	return numbers
}

func ParseNumbers(line string) []int {
	values := strings.Split(line, ",")
	numbers := make([]int, 0)
	for _, i := range values {
		v, _ := strconv.Atoi(i)
		numbers = append(numbers, v)
	}
	return numbers
}

func ParseBoards(lines []string) []Board {
	boards := make([]Board, 0)
	board := make(Board, 0)
	for _, line := range lines {
		if len(line) == 0 {
			boards = append(boards, board)
			board = make(Board, 0)
		} else {
			board = append(board, ParseLine(line))
		}
	}
	boards = append(boards, board)
	return boards
}

func (board *Board) ToString() string {
	var repr string = ""
	for _, row := range *board {
		for _, boardNumber := range row {
			switch boardNumber.extracted {
			case true:
				repr += fmt.Sprintf("[%2d] ", boardNumber.value)
			case false:
				repr += fmt.Sprintf("%4d ", boardNumber.value)
			}
		}
		repr += "\n"
	}
	return repr
}

func (board *Board) Play(n int) {
	for idRow, row := range *board {
		for idCol, boardNumber := range row {
			if boardNumber.value == n {
				(*board)[idRow][idCol].extracted = true
			}
		}
	}
}

func (board *Board) Win() bool {
	rowsExtracted := make([]int, 5)
	colsExtracted := make([]int, 5)
	for idRow, row := range *board {
		for idCol, boardNumber := range row {
			if boardNumber.extracted {
				rowsExtracted[idRow] += 1
				colsExtracted[idCol] += 1
				if rowsExtracted[idRow] == 5 || colsExtracted[idCol] == 5 {
					return true
				}
			}
		}
	}
	return false
}

func (board *Board) Score(n int) int {
	var sum int
	for _, row := range *board {
		for _, boardNumber := range row {
			if !boardNumber.extracted {
				sum += boardNumber.value
			}
		}
	}
	return sum * n
}

func Play(boards []Board, numbers []int) (int, int) {
	for _, n := range numbers {
		for boardId, board := range boards {
			board.Play(n)
			if board.Win() {
				return boardId, board.Score(n)
			}
		}
	}
	return -1, -1
}

func PlayToLose(boards []Board, numbers []int) (int, int) {
	winningBoardsMap := make(map[int]bool)
	winningBoards := make([]int, 0)
	for _, n := range numbers {
		for boardId, board := range boards {
			if winningBoardsMap[boardId] {
				continue
			}
			board.Play(n)
			if board.Win() {
				winningBoardsMap[boardId] = true
				winningBoards = append(winningBoards, boardId)
			}
			if len(winningBoards) == len(boards) {
				lastWinning := winningBoards[len(winningBoards)-1]
				return lastWinning, boards[lastWinning].Score(n)
			}
		}
	}
	return -1, -1
}

func Part1() {
	lines, _ := utils.ReadLines("input.txt")
	var numbers = ParseNumbers(lines[0])
	boards := ParseBoards(lines[2:])
	var winner, score = Play(boards, numbers)
	fmt.Printf("Part 1: Winner %d, Score %d\n", winner, score)
}

func Part2() {
	lines, _ := utils.ReadLines("input.txt")
	var numbers = ParseNumbers(lines[0])
	boards := ParseBoards(lines[2:])
	var winner, score = PlayToLose(boards, numbers)
	fmt.Printf("Part 2: Last Winner %d, Score %d\n", winner+1, score)
}

func main() {
	Part1()
}
