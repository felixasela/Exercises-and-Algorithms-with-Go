// Minesweeper

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	rows    = 5
	columns = 5
	mines   = 5
)

func main() {
	rand.Seed(time.Now().UnixNano())

	board := createBoard()
	placeMines(board)

	fmt.Println("¡Bienvenido al Buscaminas!")
	displayBoard(board)

	for {
		row, col := getUserMove()
		if isMine(board, row, col) {
			fmt.Println("¡Has perdido! Encontraste una mina.")
			displayBoardWithMines(board, row, col)
			break
		}

		revealCell(board, row, col)

		if isGameWon(board) {
			fmt.Println("¡Felicidades! Has ganado el Buscaminas.")
			break
		}

		displayBoard(board)
	}
}

func createBoard() [][]int {
	board := make([][]int, rows)
	for i := range board {
		board[i] = make([]int, columns)
	}
	return board
}

func placeMines(board [][]int) {
	for i := 0; i < mines; i++ {
		for {
			row := rand.Intn(rows)
			col := rand.Intn(columns)

			if board[row][col] == 0 {
				board[row][col] = -1
				break
			}
		}
	}
}

func displayBoard(board [][]int) {
	fmt.Println("Tablero:")
	for _, row := range board {
		for _, cell := range row {
			if cell == -1 {
				fmt.Print(". ")
			} else if cell == 0 {
				fmt.Print(". ")
			} else {
				fmt.Printf("%d ", cell)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func displayBoardWithMines(board [][]int, revealedRow, revealedCol int) {
	fmt.Println("Tablero con Minas:")
	for i, row := range board {
		for j, cell := range row {
			if cell == -1 && (i != revealedRow || j != revealedCol) {
				fmt.Print(". ")
			} else if cell == -1 {
				fmt.Print("* ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func getUserMove() (int, int) {
	var row, col int
	fmt.Print("Ingrese la fila (0-4): ")
	fmt.Scan(&row)
	fmt.Print("Ingrese la columna (0-4): ")
	fmt.Scan(&col)
	return row, col
}

func isMine(board [][]int, row, col int) bool {
	return board[row][col] == -1
}

func isGameWon(board [][]int) bool {
	for _, row := range board {
		for _, cell := range row {
			if cell >= 0 && cell <= 8 {
				return false
			}
		}
	}
	return true
}

func countAdjacentMines(board [][]int, row, col int) int {
	count := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			newRow, newCol := row+i, col+j
			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < columns && board[newRow][newCol] == -1 {
				count++
			}
		}
	}

	return count
}

func revealCell(board [][]int, row, col int) {
	if board[row][col] == 0 {
		adjacentMines := countAdjacentMines(board, row, col)
		board[row][col] = adjacentMines
		if adjacentMines == 0 {
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					newRow, newCol := row+i, col+j
					if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < columns && board[newRow][newCol] == 0 {
						revealCell(board, newRow, newCol)
					}
				}
			}
		}
	}
}
