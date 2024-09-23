package main

import "fmt"

//Check if the soduku is valid 
func isValidSudoku(board [][]byte) bool {
	// Maps to track the seen digits for rows, columns, and 3x3 sub-boxes
	rowMap := make([]map[byte]bool, 9)
	colMap := make([]map[byte]bool, 9)
	boxMap := make([]map[byte]bool, 9)

	// Initialize maps
	for i := 0; i < 9; i++ {
		rowMap[i] = make(map[byte]bool)
		colMap[i] = make(map[byte]bool)
		boxMap[i] = make(map[byte]bool)
	}

	// Iterate through the board
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := board[i][j]

			// Ignore empty cells
			if num == '.' {
				continue
			}

			// Calculate box index (0 to 8) for the 3x3 sub-box
			//i/3 and j/3 give the row and column of the sub-box in the 3x3 grid
			//and multiplying by 3 shifts it to the right box.
			boxIdx := (i/3)*3 + j/3

			// Check if the number is already seen in the row, column, or box
			if rowMap[i][num] || colMap[j][num] || boxMap[boxIdx][num] {
				return false
			}

			// Mark the number as seen in the row, column, and box
			rowMap[i][num] = true
			colMap[j][num] = true
			boxMap[boxIdx][num] = true
		}
	}

	return true
}

func main() {
	board1 := [][]byte{
		{'1', '2', '.', '.', '3', '.', '.', '.', '.'},
		{'4', '.', '.', '5', '.', '.', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '.', '3'},
		{'5', '.', '.', '.', '6', '.', '.', '.', '4'},
		{'.', '.', '.', '8', '.', '3', '.', '.', '5'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '.', '.', '.', '.', '.', '2', '.', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '8'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(board1))

	board2 := [][]byte{
		{'1', '2', '.', '.', '3', '.', '.', '.', '.'},
		{'4', '.', '.', '5', '.', '.', '.', '.', '.'},
		{'.', '9', '1', '.', '.', '.', '.', '.', '3'},
		{'5', '.', '.', '.', '6', '.', '.', '.', '4'},
		{'.', '.', '.', '8', '.', '3', '.', '.', '5'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '.', '.', '.', '.', '.', '2', '.', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '8'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(board2))
}
