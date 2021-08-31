//Implementation of back-tracking sudoku solver working on a sudoku generator in go. 

package main

import "fmt"
import "practice/packages/generators"

func main() {
	
	fmt.Printf(SudokuGenerator.GenerateSudoku()); //generates and solves a sudoku and returns a formatted string of the grid or an invalid message.

	return;
}