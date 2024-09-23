package main

import "fmt"

//Solve by bitwise operation
func isValidSudokuByBitwise(board [][]byte) bool {
	// Arrays to track the seen digits for rows, columns, and boxes
	var rowMask, colMask, boxMask [9]int

	// Iterate through the board
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := board[i][j]

			// Ignore empty cells
			if num == '.' {
				continue
			}

			// Convert the byte to a bit position (1 to 9 -> 0 to 8)
			pos := num - '1'

			// Calculate box index (0 to 8) for the 3x3 sub-box
			boxIdx := (i/3)*3 + j/3

			// Check if the number is already seen in the row, column, or box using bitmasks
			if (rowMask[i]>>pos)&1 == 1 || (colMask[j]>>pos)&1 == 1 || (boxMask[boxIdx]>>pos)&1 == 1 {
				return false
			}

			// Mark the number as seen in the row, column, and box
			rowMask[i] |= 1 << pos
			colMask[j] |= 1 << pos
			boxMask[boxIdx] |= 1 << pos
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
	fmt.Println(isValidSudokuByBitwise(board1))

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
	fmt.Println(isValidSudokuByBitwise(board2))
}
