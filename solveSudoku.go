//Implementation of back-tracking sudoku solver in go.

package main

import "fmt"
import "practice/packages/solvers"

func main() {
	
	board := [9][9]int {{ 0, 1, 4, 0, 8, 0, 2, 0, 0},
						{ 0, 0, 5, 1, 4, 0, 0, 0, 0},
						{ 0, 0, 0, 0, 0, 0, 0, 0, 0},
						{ 0, 8, 0, 0, 0, 0, 0, 0, 0},
						{ 0, 0, 0, 0, 0, 0, 0, 0, 0},
						{ 0, 3, 0, 0, 0, 0, 0, 0, 0},
						{ 0, 0, 0, 0, 0, 0, 8, 0, 0},
						{ 0, 0, 0, 0, 0, 0, 0, 0, 0},
						{ 0, 0, 0, 0, 0, 0, 0, 0, 0}}

	fmt.Printf(SudokuSolver.SolveSudoku(board)); //generates and solves a sudoku and returns a formatted string of the grid or an invalid message.

	return;
}