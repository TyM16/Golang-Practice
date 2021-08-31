package SudokuSolver;

import "practice/packages/toStrings"

//Remember that because of the way 2d arrays work, y is the first degree array and x is the second!

type coords struct{ //Build a struct to hold current coordinates.
	x, y int; 
}

const gridSize int = 9; //Grid size does not change, sudoku grid is 9x9 

//Solver Section----------------------------------------------------------------------------------------------------------------------

func SolveSudoku(grid [gridSize][gridSize]int) string {

	
	iter := coords{0, 0}; //iterator to work with coordinates. send (0, 0) to start.

//Send to the rule checker to see if sudoku is solvable and solve if so. Else print not valid. 

	if !checkInitialNumsOK(&grid, iter) { //Initial check on the grid to make sure the numbers are not invalid at the start.
		return "Bad grid!"
	} 

	if checkGrid(&grid, iter){//Return the results or no valid solution statement.
		return SudokuFormatter.FormatSudoku(&grid);
	} else {
		return "No valid solution found!";
	};
}

//Rule Checking Section--------------------------------------------------------------------------------------------------------

//Check each row/column/block for duplicate number. If found, backtrack recursively.
func checkGrid (grid *[gridSize][gridSize]int, coordinates coords) bool { //Params: 9x9 sudoku grid. Returns: true or false if grid is valid.
	if onGrid(coordinates) == false {  //If we get a complete board (step off because all cells have been filled), this will return true to let us know all cells are filled and valid.
		return true;
	}

	if isEmpty(grid, coordinates) { //If the cell is empty, continue.
		for trialNum := 1; trialNum <= 9; trialNum++ { //Try each number in loop, start at 1 and end with 9.
			if (validRow(grid, coordinates, trialNum) && validCol(grid, coordinates, trialNum) && validBlock(grid, coordinates, trialNum)){ //if all checks return true, place the number and move to next cell.
				grid[coordinates.y][coordinates.x] = trialNum;
				if(checkGrid(grid, getNextCell(coordinates))) { //Recursion here. Get next cell and recall this function. When we get a true back, return true.
					return true;
				};
				grid[coordinates.y][coordinates.x] = 0; //Reset the number to 0. 
			}
		} 

	} else {
		if(checkGrid(grid, getNextCell(coordinates))){ //Cell already taken, treat same as if we put a valid number here.
			return true;
		}
	}


	return false; //No valid numbers were found for this path, return false and go back to previous step.
}

func validRow(grid *[gridSize][gridSize]int, coordinates coords, num int) bool { //Params: 9x9 sudoku grid, coords of cell. Return false if cell conflicts with row
	for i := 0; i < gridSize; i++ { //iterate through the row to check other numbers against the one we are testing.
		if grid[i][coordinates.x] == num {
			return false; // If same number is already in row, return false because row is invalid.
		}
	}

	return true; //Otherwise, return true for valid row.
}

func validCol(grid *[gridSize][gridSize]int, coordinates coords, num int) bool { //Params: 9x9 sudoku grid, coords of cell. Return false if cell conflicts with column
	for i := 0; i < gridSize; i++ { //iterate through the row to check other numbers against the one we are testing.
		if grid[coordinates.y][i] == num {
			return false; // If same number is already in row, return false because row is invalid.
		}
	}

	return true; //Otherwise, return true for valid row.
}

func validBlock(grid *[gridSize][gridSize]int, coordinates coords, num int) bool { //Params: 9x9 sudoku grid, coords of cell. Return false if cell conflicts with block
	var xBlockOffset int = coordinates.x % 3; //Get the block of the x coordinate and calculate offset (can be 0 1 or 2)
	var yBlockOffset int = coordinates.y % 3; //Repeat for y.

	for i := 0; i < 3; i++ { //Iterate through a 3x3 block and use the offsets to determine which one
		for j := 0; j < 3; j++ {
			if grid[coordinates.y + i - yBlockOffset][coordinates.x + j - xBlockOffset] == num { //If we find our number, then the number is invalid for the cell we want to put it in and we return false.
				return false;
			}
		}
	}

	return true; //Otherwise, it's fine and we return true.
}

func isEmpty(grid *[gridSize][gridSize]int, coordinates coords) bool { //Params: 9x9 sudoku grid, coords of cell. Return true if cell is empty
	if grid[coordinates.y][coordinates.x] != 0 { //If the cell is not 0, it's not an empty cell
		return false; //return false for non-empty cells.
	}

	return true; //Otherwise it is empty.
}

//Utilities Section--------------------------------------------------------------------------------------------------------------
func checkInitialNumsOK(grid *[gridSize][gridSize]int, iter coords) bool { //Paramas 9x9 sudoku grid
	tempNum := 0; //Use this to hold the current cell's number to swap back in. This is so the check logic doesn't have to be changed above and we can change each cell to 0 before checking.
	
	for y, row := range grid { //for each row in the grid...
		iter.y = y;
		for x, _:= range row { //...and each cell in the row...
			iter.x = x;
			if (grid[y][x] != 0) {//..If it's not 0 (empty)...
				tempNum = grid[y][x];
				grid[y][x] = 0;//Set the cell to empty temporarily.
				 if !(validRow(grid, iter, tempNum) && validCol(grid, iter, tempNum) && validBlock(grid, iter, tempNum)) { //...If the number is invalid, immediately return false
				 	return false; 
				 } 
	
				grid[y][x] = tempNum; //Return the original value to the cell after each use of the temperary placeholder!
			}	
		}
	}
	
	return true; //Else, return true if the initial board is OK.
}

func onGrid(coordinates coords) bool { //Check if the coordinates are on the grid
	if (coordinates.x <= 8 && coordinates.y <=8 && coordinates.x >=0 && coordinates.y >= 0)  { //return true if on grid.
		return true;
	}

	return false; //else return false.
}

func getNextCell(coordinates coords) coords{ //Params: coordinates to increment. Get the next cell depending on where we are now.
	if coordinates.x < 8 {//Shift right if we are not at the end of the row...
		coordinates.x++
	} else { //...else shift down and reset x to the beginning of the row.
		coordinates.y++; 
		coordinates.x = 0;
	} 

	return coordinates; //Note we can shift too far if we are at the last cell, this is ok because it indicates we are done to the calling function.
}