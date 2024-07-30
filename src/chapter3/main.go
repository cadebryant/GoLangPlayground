package main

import (
	"fmt"
	"strconv"
)

type GameBoard []([]bool)
type GameBoard2 []([]int)

func main() {
	numRows := 6
	numCols := 6
	board := InitializeBoard(numRows, numCols)
	fmt.Println(board)
}

func PlayAutomaton(initialBoard GameBoard2, numGens int, neighborhoodType string, rules map[string]int) []GameBoard2 {
	boards := make([]GameBoard2, numGens+1)
	boards[0] = initialBoard
	for i := 1; i <= numGens; i++ {
		boards[i] = UpdateBoard2(boards[i-1], neighborhoodType, rules)
	}
	return boards
}

func UpdateBoard2(currBoard GameBoard2, neighborhoodType string, ruleStrings map[string]int) GameBoard2 {
	numRows := CountRows2(currBoard)
	numCols := CountCols2(currBoard)
	newBoard := InitializeBoard2(numRows, numCols)
	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			newBoard[r][c] = UpdateCell2(currBoard, r, c, neighborhoodType, ruleStrings)
		}
	}
	return newBoard
}

func UpdateCell2(currentBoard GameBoard2, r, c int, neighborhoodType string, rules map[string]int) int {
	neighborhood := NeighborhoodToString(currentBoard, r, c, neighborhoodType)
	return rules[neighborhood]
}

func NeighborhoodToString(currentBoard GameBoard2, row, col int, neighborhoodType string) string {
	neighborhood := strconv.Itoa(currentBoard[row][col])
	var offsets [][]int
	if neighborhoodType == "Moore" {
		offsets = [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}
	} else if neighborhoodType == "vonNeumann" {
		offsets = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	}
	len := len(offsets)
	for i := 0; i < len; i++ {
		x := offsets[i][0]
		y := offsets[i][1]
		if InField2(currentBoard, row+x, col+y) {
			neighborhood += strconv.Itoa(currentBoard[row+x][col+y])
		} else {
			neighborhood += "0"
		}
	}
	return neighborhood
}

func PlayGameOfLife(initialBoard GameBoard, numGens int) []GameBoard {
	boards := make([]GameBoard, numGens+1)
	boards[0] = initialBoard
	for i := 1; i <= numGens; i++ {
		boards[i] = UpdateBoard(boards[i-1])
	}
	return boards
}

func UpdateBoard(currBoard GameBoard) GameBoard {
	numRows := CountRows(currBoard)
	numCols := CountCols(currBoard)
	board := InitializeBoard(numRows, numCols)
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			board[i][j] = UpdateCell(currBoard, i, j)
		}
	}
	return board
}

func UpdateCell(board GameBoard, r, c int) bool {
	numNeighbors := CountLiveNbrs(board, r, c)
	if board[r][c] {
		return numNeighbors == 2 || numNeighbors == 3
	}
	return numNeighbors == 3
}

func InitializeBoard(numRows, numCols int) GameBoard {
	board := make(GameBoard, numRows)
	for i := range board {
		board[i] = make([]bool, numCols)
	}
	return board
}

func InitializeBoard2(numRows, numCols int) GameBoard2 {
	board := make(GameBoard2, numRows)
	for i := range board {
		board[i] = make([]int, numCols)
	}
	return board
}

func CountLiveNbrs(board GameBoard, r, c int) int {
	count := 0
	for i := r - 1; i <= r+1; i++ {
		for j := c - 1; j <= c+1; j++ {
			if (i != r || j != c) && InField(board, i, j) {
				if board[i][j] {
					count++
				}
			}
		}
	}
	return count
}

func InField2(board GameBoard2, i, j int) bool {
	numRows := CountRows2(board)
	numCols := CountCols2(board)
	return i >= 0 && i < numRows && j >= 0 && j < numCols
}

func InField(board GameBoard, i, j int) bool {
	numRows := CountRows(board)
	numCols := CountCols(board)
	return i >= 0 && i < numRows && j >= 0 && j < numCols
}

func CountRows2(board GameBoard2) int {
	return len(board)
}

func CountCols2(board GameBoard2) int {
	// assume that we have a rectangular board
	if CountRows2(board) == 0 {
		panic("Error: empty board given to CountCols")
	}
	// give # of elements in 0-th row
	return len(board[0])
}

func CountRows(board GameBoard) int {
	return len(board)
}

func CountCols(board GameBoard) int {
	// assume that we have a rectangular board
	if CountRows(board) == 0 {
		panic("Error: empty board given to CountCols")
	}
	// give # of elements in 0-th row
	return len(board[0])
}
