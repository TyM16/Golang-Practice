package SudokuFormatter;

import "strconv" //We need this to change from int values to string!

const gridSize int = 9; //Grid size does not change, sudoku grid is 9x9

func FormatSudoku (grid *[gridSize][gridSize]int) string { //Params: 9x9 sudoku .
	
	var formattedGrid string = "";
	
	for idx, val := range *grid { //Special for loop over range, value of current cell is stored in val. Remember, these are arrays of rows, we need each value inside!
		
		if idx % 3 == 0  && idx != 0 { //Add dash symbols to split the grid into 9 blocks, like a sudoku grid usually has.
			formattedGrid += "\n";

			for i := 0; i < 11; i++ { 
				formattedGrid += "- "; 
			}
		} 
		
		if idx != 0 {
			formattedGrid += "\n" //Add a newline for each new row.
		}

		for idx, val2 := range val { //Get the value of each number in the row.
			if idx % 3 == 0 && idx != 0 { formattedGrid += "| ";} //Add pipe symbols to split the grid into 9 blocks, like a sudoku grid usually has.
			formattedGrid += strconv.Itoa(val2) + " ";  //Add each number in row to string then move to next until done.
		}
	}
	
	return formattedGrid; //Return the completely formatted grid as a string.
}