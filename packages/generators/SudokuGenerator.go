package SudokuGenerator

import "practice/packages/solvers"
import "math/rand" //We need this to generate random numbers.
import "time"

const gridSize int = 9; //Grid size does not change, sudoku grid is 9x9

func GenerateSudoku() string {

	rand.Seed(time.Now().UnixNano());

	var grid [gridSize][gridSize]int;

	var result string;

	for done := false; !done; {

		var n int;

		for y, row := range grid {
			for x, _ := range row {
				n = rand.Intn(10);
				if n == 0 { 
					grid[y][x] = rand.Intn(10);
				} else {
					grid[y][x] = 0;
				}
			}
		}
		
		result = SudokuSolver.SolveSudoku(grid);

		if !(result == "Bad grid!" || result == "No valid solution found!") { //Once we get a good grid, break out of the loop. 			
			done = true;
		}

		for y, row := range grid {
			for x, _ := range row {
				grid[y][x] = 0;
			}
		}
	}

	return result;
}